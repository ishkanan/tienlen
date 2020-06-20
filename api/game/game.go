package game

import (
	"reflect"
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/ishkanan/tienlen/api/utils"
)

type connState int
type gameState int

const (
	connStateNew     connState = 1
	connStateDead    connState = 2
	gameStateInLobby gameState = 1
	gameStateRunning gameState = 2
	gameStatePaused  gameState = 3
)

type context struct {
	Player     *player
	Connection IMessageSink
}

// Game encapsulates the core game logic and high-level comms to players
type Game struct {
	players     players
	state       gameState
	lastPlayed  []card
	firstRound  bool
	newRound    bool
	connections map[string]context
	mutex       *sync.Mutex
}

// NewGame builds a new game instance and calls Init()
func NewGame() *Game {
	g := &Game{}
	g.Init()
	return g
}

// Init initialises the game instance - do not call this if NewGame was used
func (g *Game) Init() {
	g.players = make(players, 0, 4)
	g.state = gameStateInLobby
	g.lastPlayed = nil
	g.connections = map[string]context{}
	g.mutex = &sync.Mutex{}
}

// IsAcceptingConnections indicates if the game can accept more player connections
func (g *Game) IsAcceptingConnections() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	totalConnections := len(g.players) - g.disconnectedCount() + g.unmappedCount()
	return (g.state == gameStateInLobby && totalConnections < 4) || (g.state == gameStatePaused && totalConnections < len(g.players))
}

// ConnectionStateChanged informs the game of a new or expired player connection
func (g *Game) ConnectionStateChanged(connUUID uuid.UUID, conn IMessageSink, state connState) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	connID := connUUID.String()

	if state == connStateNew {
		g.connections[connID] = context{Connection: conn}
		return
	}

	// disconnection

	player := g.connections[connID].Player
	if player != nil {
		player.Connected = false
		g.sendToAllPlayers(playerDisconnectedResponse{Player: *player})
		utils.LogInfo("ConnectionStateChanged: %s has disconnected", player.Name)
		if g.state == gameStateInLobby {
			// no need to keep place for player if game hasn't started
			g.players = g.players.DeleteByID(player.ID)
			// but we re-number positions to be sequential
			for _, p := range g.players {
				if p.Position > player.Position {
					p.Position--
				}
			}
		} else if g.state == gameStateRunning {
			g.state = gameStatePaused
			g.sendToAllPlayers(gamePausedResponse{})
			utils.LogInfo("ConnectionStateChanged: Game is paused due to player disconnect")
		}
	} else {
		utils.LogDebug("ProcessRequest: %s has disconnected", connID)
	}
	delete(g.connections, connID)

	if len(g.players) == g.disconnectedCount()+g.unmappedCount() {
		g.Init()
		utils.LogInfo("ConnectionStateChanged: All players have left, game is reset")
	}
	g.sendStateToAllPlayers()
}

// ProcessRequest informs the game about a request received over a player connection
func (g *Game) ProcessRequest(connUUID uuid.UUID, request interface{}, requestType reflect.Type) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	connID := connUUID.String()

	if requestType == reflect.TypeOf(joinGameRequest{}) {
		req := request.(joinGameRequest)
		g.processJoinGameRequest(connID, req)
		return
	}

	if g.connections[connID].Player == nil {
		utils.LogDebug("ProcessRequest: request ignored for %s - %+v", connID, request)
		g.sendOnConnection(connID, errorResponse{Kind: errKindNotAuthorised})
		return
	}

	if requestType == reflect.TypeOf(startGameRequest{}) {
		g.processStartGameRequest(connID)
		return
	}

	if g.state != gameStateRunning {
		g.sendOnConnection(connID, errorResponse{Kind: errKindNotAuthorised})
		return
	}

	if g.players.CurrentTurn().ID != g.connections[connID].Player.ID {
		g.sendOnConnection(connID, errorResponse{Kind: errKindOutOfTurn})
		return
	}

	if requestType == reflect.TypeOf(turnPassRequest{}) {
		g.processTurnPassRequest(connID)
		return
	}

	if requestType == reflect.TypeOf(turnPlayRequest{}) {
		req := request.(turnPlayRequest)
		g.processTurnPlayRequest(connID, req)
	}
}

func (g *Game) processJoinGameRequest(connID string, req joinGameRequest) {
	thePlayer := g.players.GetByName(req.PlayerName)

	rejoined := thePlayer != nil
	for cID, context := range g.connections {
		// prevent someone hijacking a connected player
		if cID != connID && context.Player != nil {
			rejoined = rejoined && context.Player.Name != req.PlayerName
		}
	}

	if !rejoined {
		if g.state == gameStatePaused {
			g.sendOnConnection(connID, errorResponse{Kind: errKindGameFull})
			g.connections[connID].Connection.Close()
			utils.LogInfo("processJoinGameRequest: %s tried to join, but game is full", connID)
			return
		}
		id, _ := uuid.Parse(connID)
		position := g.players.NextAvailablePosition()
		name := strings.TrimSpace(req.PlayerName)
		if g.players.GetByName(name) != nil {
			g.sendOnConnection(connID, errorResponse{Kind: errKindNameTaken})
			g.connections[connID].Connection.Close()
			utils.LogInfo("processJoinGameRequest: %s tried to use name %s, but it's taken", connID, name)
			return
		}
		if name == "" {
			name = g.players.theyWhoNotBeNamed()
		}
		thePlayer = &player{
			ID:       id,
			Name:     name,
			Position: position,
		}
		g.players = append(g.players, thePlayer)
	}

	thePlayer.Connected = true
	context := g.connections[connID]
	context.Player = thePlayer
	g.connections[connID] = context

	g.sendToAllPlayers(playerJoinedResponse{Player: *thePlayer})
	utils.LogInfo("processJoinGameRequest: %s has joined the game on %s", thePlayer.Name, connID)

	if rejoined && g.disconnectedCount() == 0 {
		g.state = gameStateRunning
		g.sendToAllPlayers(gameResumedResponse{})
		utils.LogInfo("processJoinGameRequest: All players have re-joined, game is resumed")
	}

	g.sendStateToAllPlayers()
}

func (g *Game) processStartGameRequest(connID string) {
	thePlayer := g.connections[connID].Player

	if g.state != gameStateInLobby || len(g.players) < 2 || g.disconnectedCount() > 0 || g.unmappedCount() > 0 {
		g.sendOnConnection(connID, errorResponse{Kind: errKindNotAuthorised})
		utils.LogDebug("processStartGameRequest: Unauthorised attempt by %s", thePlayer.Name)
		return
	}

	deck := buildShuffledDeck()
	for i, player := range g.players {
		player.IsPassed = false
		player.IsTurn = false
		player.Hand = deck[i*13 : (i*13)+13]
		player.CardsLeft = 13
	}

	first := g.players.WonLastGame()
	if first == nil {
		first = g.players.WithLowestCard()
	}
	g.lastPlayed = nil
	first.IsTurn = true
	g.state = gameStateRunning
	g.firstRound = true
	g.newRound = true

	g.sendStateToAllPlayers()
	g.sendToAllPlayers(gameStartedResponse{Player: *thePlayer})
	utils.LogInfo("processStartGameRequest: %s has started the game, %s starts play", thePlayer.Name, first.Name)
}

func (g *Game) processTurnPassRequest(connID string) {
	thePlayer := g.connections[connID].Player

	if g.newRound {
		g.sendOnConnection(connID, errorResponse{Kind: errKindMustPlay})
		utils.LogDebug("processTurnPassRequest: Unauthorised attempt by %s", thePlayer.Name)
		return
	}

	thePlayer.IsPassed = true
	thePlayer.IsTurn = false
	nextPlayer := g.players.NextTurn(thePlayer)
	nextPlayer.IsTurn = true

	// if only one player left who hasn't passed, start new round
	if g.players.NextTurn(nextPlayer).ID == nextPlayer.ID {
		for _, player := range g.players {
			player.IsPassed = false
		}
		g.lastPlayed = nil
		g.newRound = true
	}

	g.sendStateToAllPlayers()
	g.sendToAllPlayers(turnPassedResponse{Player: *thePlayer})
	if g.newRound {
		utils.LogInfo("processTurnPassRequest: %s has won the round", nextPlayer.Name)
		g.sendToAllPlayers(roundWonResponse{Player: *nextPlayer})
	}
}

func (g *Game) processTurnPlayRequest(connID string, req turnPlayRequest) {
	thePlayer := g.connections[connID].Player

	if len(req.Cards) == 0 || len(req.Cards) > len(thePlayer.Hand) {
		g.sendOnConnection(connID, errorResponse{Kind: errKindInvalidCards})
		utils.LogDebug("processTurnPlayRequest: Rejected proposed cards from %s - invalid cards", thePlayer.Name)
		return
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
	} else if !g.newRound && !areBetterCardsThan(cardsToPlay, g.lastPlayed) {
		err = errKindCardsNotBetter
	} else if g.firstRound && !thePlayer.WonLastGame && cardInSet(lowestCard.GlobalRank, cardsToPlay) == -1 {
		err = errKindMustPlayLowest
	}
	if err != errKindLobbyNotReady {
		msg := map[errorKind]string{
			errKindInvalidCards:   "invalid cards",
			errKindInvalidPattern: "invalid pattern",
			errKindCardsNotBetter: "cards not better than last played",
			errKindMustPlayLowest: "must played lowest",
		}[err]
		g.sendOnConnection(connID, errorResponse{Kind: err})
		utils.LogDebug("processTurnPlayRequest: Rejected proposed cards from %s - %s", thePlayer.Name, msg)
		return
	}

	g.lastPlayed = cardsToPlay
	g.firstRound = false
	g.newRound = false
	thePlayer.CardsLeft = len(newHand)
	thePlayer.Hand = newHand
	thePlayer.IsTurn = false
	g.players.NextTurn(thePlayer).IsTurn = len(newHand) > 0
	won := len(thePlayer.Hand) == 0 || (determinePattern(cardsToPlay) == patternQuad && cardsToPlay[0].FaceValue == 2)
	if won {
		g.state = gameStateInLobby
		for _, player := range g.players {
			player.CardsLeft = 0
			player.Hand = nil
			player.IsPassed = false
			player.IsTurn = false
			player.WonLastGame = false
		}
		thePlayer.WonLastGame = true
	}

	g.sendStateToAllPlayers()
	g.sendToAllPlayers(turnPlayedResponse{
		Player: *thePlayer,
		Cards:  cardsToPlay,
	})
	utils.LogInfo("processTurnPlayRequest: %s played %+v", thePlayer.Name, cardsToPlay)

	if won {
		g.sendToAllPlayers(gameWonResponse{Player: *thePlayer})
		utils.LogInfo("processTurnPlayRequest: %s has won the game", thePlayer.Name)
	}
}

// returns the number of disconnected players (who we've kept places for)
func (g Game) disconnectedCount() int {
	count := 0
	for _, player := range g.players {
		found := false
		for _, context := range g.connections {
			found = found || (context.Player != nil && context.Player.ID == player.ID)
		}
		if !found {
			count++
		}
	}
	return count
}

// returns the number of connections that are not yet mapped to players
func (g Game) unmappedCount() int {
	count := 0
	for _, context := range g.connections {
		if context.Player == nil {
			count++
		}
	}
	return count
}

// sends a response-type message to all players
func (g Game) sendToAllPlayers(response interface{}) {
	for _, context := range g.connections {
		if context.Player != nil {
			context.Connection.Send(response)
		}
	}
}

// sends a response-type message to a connection (mapped or un-mapped)
func (g Game) sendOnConnection(connID string, response interface{}) {
	g.connections[connID].Connection.Send(response)
}

// sends a customised gameStateRefreshResponse to each player
func (g Game) sendStateToAllPlayers() {
	for _, context := range g.connections {
		opponents := []player{}
		for _, player := range g.players {
			if player.ID != context.Player.ID {
				opponents = append(opponents, *player)
			}
		}

		context.Connection.Send(gameStateRefreshResponse{
			Opponents:  opponents,
			Self:       *context.Player,
			SelfHand:   context.Player.Hand,
			GameState:  g.state,
			LastPlayed: g.lastPlayed,
			FirstRound: g.firstRound,
			NewRound:   g.newRound,
		})
	}
}
