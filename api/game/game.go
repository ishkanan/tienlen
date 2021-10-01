package game

import (
	"reflect"

	"github.com/google/uuid"

	"github.com/ishkanan/tienlen/api/utils"
)

type connState int
type GameState int

const (
	connStateNew     connState = 1
	connStateDead    connState = 2
	GameStateInLobby GameState = 1
	GameStateRunning GameState = 2
	GameStatePaused  GameState = 3
	maxNameLength              = 35
)

// PlayerContext links a player with a connection
type PlayerContext struct {
	Player     *player
	Connection IMessageSink
}

// GameEngineState encapsulates all state of a game
type GameEngineState struct {
	Players             players
	GameState           GameState
	LastPlayedHand      []card
	IsFirstRound        bool
	IsNewRound          bool
	Connections         map[string]*PlayerContext
	WinPlaces           players
	PlayerPlacedInRound bool
}

// GameEngine provides a namespace for all game engine functions
// and implements the IMessageSource and ISender interfaces
type GameEngine struct{}

// NewGameEngineState builds a new, fully reset game state
func NewGameEngineState() GameEngineState {
	return GameEngineState{
		Players:     make(players, 0, 4),
		GameState:   GameStateInLobby,
		Connections: map[string]*PlayerContext{},
		WinPlaces:   make(players, 0, 3),
	}
}

// ResetGameEngineState fully resets all game state, including connections
func ResetGameEngineState(state GameEngineState) GameEngineState {
	state.Players = make(players, 0, 4)
	state.GameState = GameStateInLobby
	state.LastPlayedHand = nil
	state.Connections = map[string]*PlayerContext{}
	state.WinPlaces = make(players, 0, 3)
	state.PlayerPlacedInRound = false
	return state
}

// NewGameEngine builds a new game engine
func NewGameEngine() GameEngine {
	return GameEngine{}
}

// DisconnectedPlayerCount returns the number of disconnected players for whom places are reserved
func (g GameEngine) DisconnectedPlayerCount(state GameEngineState) int {
	count := 0
	for _, player := range state.Players {
		found := false
		for _, context := range state.Connections {
			found = found || (context.Player != nil && context.Player.Name == player.Name)
		}
		if !found {
			count++
		}
	}
	return count
}

// UnmappedConnectionCount returns the number of connections that are not yet mapped to players
func (g GameEngine) UnmappedConnectionCount(state GameEngineState) int {
	count := 0
	for _, context := range state.Connections {
		if context.Player == nil {
			count++
		}
	}
	return count
}

// IsAcceptingConnections indicates if the game can accept more player connections
func (g GameEngine) IsAcceptingConnections(state GameEngineState) bool {
	totalConnections := len(state.Players) - g.DisconnectedPlayerCount(state) + g.UnmappedConnectionCount(state)
	return (state.GameState == GameStateInLobby && totalConnections < 4) ||
		(state.GameState == GameStatePaused && totalConnections < len(state.Players))
}

// --------------------------- Event handlers ----------------------------------------

// ConnectionStateChanged informs the game of a new or expired player connection
func (g GameEngine) ConnectionStateChanged(state GameEngineState, connUUID uuid.UUID, conn IMessageSink, connState connState, pump ISender) GameEngineState {
	connID := connUUID.String()

	if connState == connStateNew {
		state.Connections[connID] = &PlayerContext{Connection: conn}
		return state
	}

	// disconnection

	player := state.Connections[connID].Player
	if player != nil {
		player.Connected = false
		pump.SendResponseToAllPlayers(state, playerDisconnectedResponse{Player: *player})
		utils.LogInfo("ConnectionStateChanged:: \"%s\" has disconnected", player.Name)
		if state.GameState == GameStateInLobby {
			state.WinPlaces = make(players, 0, 3)
			state.PlayerPlacedInRound = false
			state.Players.ResetScores()
			// no need to keep place for player if game hasn't started
			state.Players = state.Players.DeleteByName(player.Name)
			// but we re-number positions to be sequential
			for _, p := range state.Players {
				if p.Position > player.Position {
					p.Position--
				}
			}
		} else if state.GameState == GameStateRunning {
			state.GameState = GameStatePaused
			pump.SendResponseToAllPlayers(state, gamePausedResponse{})
			utils.LogInfo("ConnectionStateChanged:: Game is paused due to player disconnect")
		}
	} else {
		utils.LogDebug("ConnectionStateChanged:: \"%s\" has disconnected", connID)
	}
	delete(state.Connections, connID)

	if len(state.Players) == g.DisconnectedPlayerCount(state)+g.UnmappedConnectionCount(state) {
		state = ResetGameEngineState(state)
		utils.LogInfo("ConnectionStateChanged:: All players have left, game is reset")
	}
	pump.SendStateToAllPlayers(state)

	return state
}

// ProcessRequest informs the game about a request received over a player connection
func (g GameEngine) ProcessRequest(state GameEngineState, connUUID uuid.UUID, request interface{}, requestType reflect.Type, pump ISender) GameEngineState {
	connID := connUUID.String()

	if requestType == reflect.TypeOf(resetGameRequest{}) {
		return g.processResetGameRequest(state, connID, pump)
	}

	if requestType == reflect.TypeOf(joinGameRequest{}) {
		return g.processJoinGameRequest(state, connID, request.(joinGameRequest), pump)
	}

	if state.Connections[connID].Player == nil {
		utils.LogDebug("ProcessRequest: request ignored for unlinked connection %s - %+v", connID, request)
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindNotAuthorised})
		return state
	}

	if requestType == reflect.TypeOf(startGameRequest{}) {
		return g.processStartGameRequest(state, connID, pump)
	}

	if requestType == reflect.TypeOf(changeNameRequest{}) {
		return g.processChangeNameRequest(state, connID, request.(changeNameRequest), pump)
	}

	if state.GameState != GameStateRunning {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindNotAuthorised})
		return state
	}

	if state.Players.CurrentTurn().Name != state.Connections[connID].Player.Name {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindOutOfTurn})
		return state
	}

	if requestType == reflect.TypeOf(turnPassRequest{}) {
		return g.processTurnPassRequest(state, connID, pump)
	}

	if requestType == reflect.TypeOf(turnPlayRequest{}) {
		return g.processTurnPlayRequest(state, connID, request.(turnPlayRequest), pump)
	}

	return state
}

// processes a resetGameRequest message
func (g GameEngine) processResetGameRequest(state GameEngineState, connID string, pump ISender) GameEngineState {
	thePlayer := state.Connections[connID].Player

	state.GameState = GameStateInLobby
	state.IsFirstRound = true
	state.Players = state.Players.DeleteDisconnected()
	state.WinPlaces = make(players, 0, 3)
	state = g.setNewRound(state)
	state.Players.ResetAllGameStatuses()
	for _, player := range state.Players {
		player.Hand = []card{}
		player.CardsLeft = 0
	}

	pump.SendStateToAllPlayers(state)
	pump.SendResponseToAllPlayers(state, gameResetResponse{Player: *thePlayer})
	utils.LogInfo("processResetGameRequest:: \"%s\" has reset the game", thePlayer.Name)
	return state
}

// processes a joinGameRequest message
func (g GameEngine) processJoinGameRequest(state GameEngineState, connID string, req joinGameRequest, pump ISender) GameEngineState {
	thePlayer := state.Players.GetByName(req.PlayerName)

	rejoined := thePlayer != nil
	for cID, context := range state.Connections {
		// prevent someone hijacking a connected player
		if cID != connID && context.Player != nil {
			rejoined = rejoined && context.Player.Name != req.PlayerName
		}
	}

	if !rejoined {
		if state.GameState == GameStatePaused {
			pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindGameFull})
			state.Connections[connID].Connection.Close()
			utils.LogInfo("processJoinGameRequest:: %s tried to join, but game is full", connID)
			return state
		}
		position := state.Players.NextAvailablePosition()
		name := cleanPlayerName(req.PlayerName, state.Players, maxNameLength)
		if state.Players.GetByName(name) != nil {
			pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindInvalidName})
			state.Connections[connID].Connection.Close()
			utils.LogInfo("processJoinGameRequest:: %s tried to use name \"%s\", but it's taken", connID, name)
			return state
		}
		if name == "" {
			name = theyWhoNotBeNamed(state.Players, maxNameLength)
		}
		thePlayer = &player{
			Name:     name,
			Position: position,
		}
		state.Players = append(state.Players, thePlayer)
		state.Players.ResetAllGameStatuses()
		state.Players.ResetScores()
		state.WinPlaces = make(players, 0, 3)
		state.PlayerPlacedInRound = false
	}

	thePlayer.Connected = true
	state.Connections[connID].Player = thePlayer

	pump.SendResponseToAllPlayers(state, playerJoinedResponse{Player: *thePlayer})
	utils.LogInfo("processJoinGameRequest:: \"%s\" has joined the game on %s", thePlayer.Name, connID)

	if rejoined && g.DisconnectedPlayerCount(state) == 0 {
		state.GameState = GameStateRunning
		pump.SendResponseToAllPlayers(state, gameResumedResponse{})
		utils.LogInfo("processJoinGameRequest:: All players have re-joined, game is resumed")
	}

	pump.SendStateToAllPlayers(state)
	return state
}

// processes a startGameRequest message
func (g GameEngine) processStartGameRequest(state GameEngineState, connID string, pump ISender) GameEngineState {
	thePlayer := state.Connections[connID].Player

	if state.GameState != GameStateInLobby || len(state.Players) < 2 || g.DisconnectedPlayerCount(state) > 0 || g.UnmappedConnectionCount(state) > 0 {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindNotAuthorised})
		utils.LogDebug("processStartGameRequest:: Unauthorised attempt by \"%s\"", thePlayer.Name)
		return state
	}

	deck := buildShuffledDeck()
	for i, player := range state.Players {
		player.Hand = globalRankSort(deck[i*13 : (i*13)+13])
		player.CardsLeft = 13
	}

	first := state.Players.WonLastGame()
	if first == nil {
		first = state.Players.WithLowestCard()
	}
	first.IsTurn = true
	state.GameState = GameStateRunning
	state.IsFirstRound = true
	state.WinPlaces = make(players, 0, 3)
	state = g.setNewRound(state)

	pump.SendStateToAllPlayers(state)
	pump.SendResponseToAllPlayers(state, gameStartedResponse{Player: *thePlayer})
	utils.LogInfo("processStartGameRequest:: \"%s\" has started the game, \"%s\" starts play", thePlayer.Name, first.Name)
	return state
}

// processes a turnPassRequest message
func (g GameEngine) processTurnPassRequest(state GameEngineState, connID string, pump ISender) GameEngineState {
	thePlayer := state.Connections[connID].Player

	if state.IsNewRound {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindMustPlay})
		utils.LogDebug("processTurnPassRequest:: Unauthorised attempt by \"%s\"", thePlayer.Name)
		return state
	}

	utils.LogInfo("processTurnPassRequest:: \"%s\" has passed their turn", thePlayer.Name)

	thePlayer.IsPassed = true
	thePlayer.IsTurn = false
	nextPlayer := state.Players.NextTurn(thePlayer)
	nextPlayer.IsTurn = true

	if state.Players.PassedAndPlacedCount() == len(state.Players) {
		// no-one beat cards of most recent placed player, so start new round
		state = g.setNewRound(state)
		nextPlayer.IsTurn = false
		nextPlayer = state.Players.NextTurn(nextPlayer)
		nextPlayer.IsTurn = true
	} else if state.Players.NextTurn(nextPlayer).Position == nextPlayer.Position && !state.PlayerPlacedInRound {
		// only 1 player left who hasn't passed, so start new round
		state = g.setNewRound(state)
	}

	pump.SendStateToAllPlayers(state)
	pump.SendResponseToAllPlayers(state, turnPassedResponse{Player: *thePlayer})
	if state.IsNewRound {
		utils.LogInfo("processTurnPassRequest:: \"%s\" has won the round", nextPlayer.Name)
		pump.SendResponseToAllPlayers(state, roundWonResponse{Player: *nextPlayer})
	}
	return state
}

// processes a turnPlayRequest message
func (g GameEngine) processTurnPlayRequest(state GameEngineState, connID string, req turnPlayRequest, pump ISender) GameEngineState {
	thePlayer := state.Connections[connID].Player

	if len(req.Cards) == 0 || len(req.Cards) > len(thePlayer.Hand) {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindInvalidCards})
		utils.LogDebug("processTurnPlayRequest:: Rejected proposed cards from \"%s\" - invalid cards", thePlayer.Name)
		return state
	}

	cardsToPlay := make([]card, 0, len(req.Cards))
	lowestCard := globalRankSort(thePlayer.Hand)[0]
	newHand := append([]card(nil), thePlayer.Hand...)

	for _, globalRank := range req.Cards {
		i := cardInSet(globalRank, newHand)
		if i >= 0 {
			cardsToPlay = append(cardsToPlay, newHand[i])
			newHand = append(newHand[:i], newHand[i+1:]...)
		}
	}

	err := errKindLobbyNotReady
	if len(cardsToPlay) != len(req.Cards) {
		err = errKindInvalidCards
	} else if determinePattern(cardsToPlay) == patternInvalid {
		err = errKindInvalidPattern
	} else if !state.IsNewRound && !areBetterCardsThan(cardsToPlay, state.LastPlayedHand) {
		err = errKindCardsNotBetter
	} else if state.IsFirstRound && !thePlayer.WonLastGame && cardInSet(lowestCard.GlobalRank, cardsToPlay) == -1 {
		err = errKindMustPlayLowest
	}
	if err != errKindLobbyNotReady {
		msg := map[errorKind]string{
			errKindInvalidCards:   "invalid cards",
			errKindInvalidPattern: "invalid pattern",
			errKindCardsNotBetter: "cards not better than last played",
			errKindMustPlayLowest: "must play lowest",
		}[err]
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: err})
		utils.LogDebug("processTurnPlayRequest:: Rejected proposed cards from \"%s\" - %s", thePlayer.Name, msg)
		return state
	}

	thePlayer.CardsLeft = len(newHand)
	thePlayer.Hand = newHand
	thePlayer.IsTurn = false
	state.IsFirstRound = false
	state.LastPlayedHand = cardsToPlay
	state.Players.NextTurn(thePlayer).IsTurn = true
	state.Players.SetLastPlayed(thePlayer)
	state.PlayerPlacedInRound = false
	state.IsNewRound = false

	placed := len(thePlayer.Hand) == 0 || (determinePattern(cardsToPlay) == patternQuad && cardsToPlay[0].FaceValue == 2)
	if placed {
		state.WinPlaces = append(state.WinPlaces, thePlayer)
		state.PlayerPlacedInRound = true

		if len(state.WinPlaces) == len(state.Players)-1 {
			// all possible players have secured a place
			state.GameState = GameStateInLobby
			state.Players.ResetAllGameStatuses()
			state.WinPlaces[0].WonLastGame = true
			for i, player := range state.WinPlaces {
				player.Score += len(state.Players) - 1 - i
			}
		} else {
			// there are still some players to secure a place (i.e. 3-4 player game)
			if state.Players.PassedAndPlacedCount() == len(state.Players) {
				// the player who just placed was the last non-passed player,
				// so we start a new round for the next player
				state = g.setNewRound(state)
				thePlayer.IsTurn = false
				state.Players.NextTurn(thePlayer).IsTurn = true
			}
		}
	}

	if len(state.Players)-state.Players.PassedAndPlacedCount() == 1 {
		state = g.setNewRound(state)
	}

	pump.SendStateToAllPlayers(state)
	pump.SendResponseToAllPlayers(state, turnPlayedResponse{
		Player: *thePlayer,
		Cards:  cardsToPlay,
	})
	utils.LogInfo("processTurnPlayRequest:: \"%s\" played %+v", thePlayer.Name, cardsToPlay)

	if placed {
		pump.SendResponseToAllPlayers(state, playerPlacedResponse{Player: *thePlayer, Place: len(state.WinPlaces)})
		utils.LogInfo("processTurnPlayRequest:: \"%s\" has played all their cards and got %s place", thePlayer.Name, utils.Ordinal(len(state.WinPlaces)))
	} else if state.IsNewRound {
		utils.LogInfo("processTurnPlayRequest:: \"%s\" has won the round", thePlayer.Name)
		pump.SendResponseToAllPlayers(state, roundWonResponse{Player: *thePlayer})
	}

	if state.GameState == GameStateInLobby {
		pump.SendResponseToAllPlayers(state, gameWonResponse{Player: *state.WinPlaces[0]})
		utils.LogInfo("processTurnPlayRequest:: \"%s\" has won the game", state.WinPlaces[0].Name)
	}
	return state
}

// processes a changeNameRequest message
func (g GameEngine) processChangeNameRequest(state GameEngineState, connID string, req changeNameRequest, pump ISender) GameEngineState {
	thePlayer := state.Connections[connID].Player

	name := cleanPlayerName(req.PlayerName, state.Players, maxNameLength)
	if name == "" {
		pump.SendResponseOnConnection(state, connID, errorResponse{Kind: errKindInvalidName})
		return state
	}

	utils.LogInfo("processChangeNameRequest:: \"%s\" is now known as \"%s\"", thePlayer.Name, name)
	oldPlayer := *thePlayer
	thePlayer.Name = name
	pump.SendResponseToAllPlayers(state, nameChangedResponse{OldPlayer: oldPlayer, NewPlayer: *thePlayer})
	pump.SendStateToAllPlayers(state)
	return state
}

// starts a new mid-game round
func (g GameEngine) setNewRound(state GameEngineState) GameEngineState {
	state.LastPlayedHand = nil
	state.IsNewRound = true
	state.PlayerPlacedInRound = false
	state.Players.SetLastPlayed(nil)
	state.Players.UnsetPassed()
	return state
}

// --------------------------- ISender implementation --------------------------------

// SendResponseToAllPlayers sends a response-type message to all players
func (g GameEngine) SendResponseToAllPlayers(state GameEngineState, response interface{}) map[*PlayerContext]error {
	errs := map[*PlayerContext]error{}
	for _, context := range state.Connections {
		if context.Player != nil {
			if err := context.Connection.Send(response); err != nil {
				errs[context] = err
			}
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

// SendResponseOnConnection sends a response-type message to a connection (mapped or un-mapped)
func (g GameEngine) SendResponseOnConnection(state GameEngineState, connID string, response interface{}) error {
	return state.Connections[connID].Connection.Send(response)
}

// SendStateToAllPlayers sends a gameStateRefreshResponse message to each player
func (g GameEngine) SendStateToAllPlayers(state GameEngineState) map[*PlayerContext]error {
	winPlaces := make([]player, 0, 3)
	for _, player := range state.WinPlaces {
		winPlaces = append(winPlaces, *player)
	}

	errs := map[*PlayerContext]error{}

	for _, context := range state.Connections {
		if context.Player == nil {
			continue
		}

		opponents := make([]player, 0, 3)
		for _, player := range state.Players {
			if player.Name != context.Player.Name {
				opponents = append(opponents, *player)
			}
		}

		if err := context.Connection.Send(gameStateRefreshResponse{
			Opponents:  opponents,
			Self:       *context.Player,
			SelfHand:   context.Player.Hand,
			GameState:  state.GameState,
			LastPlayed: state.LastPlayedHand,
			FirstRound: state.IsFirstRound,
			NewRound:   state.IsNewRound,
			WinPlaces:  winPlaces,
		}); err != nil {
			errs[context] = err
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
