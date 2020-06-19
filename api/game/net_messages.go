package game

// Message represents a request or response that can travel over a connection
type Message struct {
	Kind string `json:"kind"`
	// Data is a base64-encoded JSON string whose decoded bytes can be marshalled into a request or response type
	Data string `json:"data"`
}

// links a new connection with a player
type joinGameRequest struct {
	PlayerName string `json:"playerName"`
}

// informs all players of a newly-connected player
type playerJoinedResponse struct {
	Player player `json:"player"`
}

// informs all players of a disconnection event
type playerDisconnectedResponse struct {
	Player player `json:"player"`
}

// starts a ready game
type startGameRequest struct{}

// informs all players that the game has started
type gameStartedResponse struct {
	Player player `json:"player"`
}

// informs all players that the game is paused
type gamePausedResponse struct{}

// informs all players that the paused game has resumed
type gameResumedResponse struct{}

// skips the current players turn
type turnPassRequest struct{}

// informs all players of the current player skipping their turn
type turnPassedResponse struct {
	Player player `json:"player"`
}

// informs all players who won the current round
type roundWonResponse struct {
	Player player `json:"player"`
}

// contains cards a player wishes to play for their current turn
type turnPlayRequest struct {
	Cards []int `json:"cards"` // ... of global rank
}

// informs all players of the cards played for the turn
type turnPlayedResponse struct {
	Player player `json:"player"`
	Cards  []card `json:"cards"`
}

// informs all players of the winner of the game
type gameWonResponse struct {
	Player player `json:"player"`
}

// provides all players with a full game state refresh
type gameStateRefreshResponse struct {
	Opponents  []player  `json:"opponents"`
	Self       player    `json:"self"`
	SelfHand   []card    `json:"selfHand"`
	GameState  gameState `json:"gameState"`
	LastPlayed []card    `json:"lastPlayed"`
	FirstRound bool      `json:"firstRound"`
	NewRound   bool      `json:"newRound"`
}

type errorKind int

const (
	errKindLobbyNotReady  errorKind = 1
	errKindNotAuthorised  errorKind = 2
	errKindOutOfTurn      errorKind = 3
	errKindMustPlay       errorKind = 4
	errKindInvalidCards   errorKind = 5
	errKindInvalidPattern errorKind = 6
	errKindCardsNotBetter errorKind = 7
	errKindMustPlayLowest errorKind = 8
	errKindNameTaken      errorKind = 9
	errKindGameFull       errorKind = 10
)

// informs a player of an invalid request
type errorResponse struct {
	Kind errorKind `json:"kind"`
}
