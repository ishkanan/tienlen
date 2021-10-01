package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type senderMock struct{}

func (s senderMock) SendResponseToAllPlayers(state GameEngineState, response interface{}) map[*PlayerContext]error {
	return nil
}

func (s senderMock) SendResponseOnConnection(state GameEngineState, connID string, response interface{}) error {
	return nil
}

func (s senderMock) SendStateToAllPlayers(state GameEngineState) map[*PlayerContext]error {
	return nil
}

func TestNewGameEngineState(t *testing.T) {
	expected := GameEngineState{
		Players:     make(players, 0, 4),
		GameState:   GameStateInLobby,
		Connections: map[string]*PlayerContext{},
		WinPlaces:   make(players, 0, 3),
	}
	state := NewGameEngineState()
	assert.Equal(t, expected, state)
}

func TestResetGameEngineState(t *testing.T) {
	expected := NewGameEngineState()
	expected.Players = make(players, 0, 4)
	expected.GameState = GameStateInLobby
	expected.LastPlayedHand = nil
	expected.Connections = map[string]*PlayerContext{}
	expected.WinPlaces = make(players, 0, 3)
	expected.PlayerPlacedInRound = false
	state := NewGameEngineState()
	state = ResetGameEngineState(state)
	assert.Equal(t, expected, state)
}

func TestDisconnectedPlayerCount(t *testing.T) {
	state := NewGameEngineState()
	engine := NewGameEngine()
	assert.Equal(t, 0, engine.DisconnectedPlayerCount(state))

	state.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
	}
	assert.Equal(t, 0, engine.DisconnectedPlayerCount(state))
	state.Connections["002"].Player = nil
	assert.Equal(t, 1, engine.DisconnectedPlayerCount(state))
}

func TestUnmappedConnectionCount(t *testing.T) {
	state := NewGameEngineState()
	engine := NewGameEngine()
	assert.Equal(t, 0, engine.UnmappedConnectionCount(state))

	state.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
	}
	assert.Equal(t, 0, engine.UnmappedConnectionCount(state))

	state.Connections["001"].Player = nil
	state.Connections["002"].Player = nil
	assert.Equal(t, 2, engine.UnmappedConnectionCount(state))
}

func TestIsAcceptingConnections(t *testing.T) {
	state := NewGameEngineState()
	engine := NewGameEngine()
	assert.Equal(t, true, engine.IsAcceptingConnections(state))

	state.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
	}
	assert.Equal(t, true, engine.IsAcceptingConnections(state))

	state.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
		{Name: "Player 4"},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
		"003": {Player: state.Players[2]},
		"004": {Player: state.Players[3]},
	}
	assert.Equal(t, false, engine.IsAcceptingConnections(state))

	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
		"003": {Player: state.Players[2]},
	}
	state.GameState = GameStatePaused
	assert.Equal(t, true, engine.IsAcceptingConnections(state))
}

func TestSetNewRound(t *testing.T) {
	state := NewGameEngineState()
	state.LastPlayedHand = []card{}
	state.IsNewRound = false
	state.PlayerPlacedInRound = true
	state.Players = players{
		{Name: "Player 1", LastPlayed: false, IsPassed: true},
		{Name: "Player 2", LastPlayed: true, IsPassed: false},
		{Name: "Player 3", LastPlayed: false, IsPassed: true},
		{Name: "Player 4", LastPlayed: false, IsPassed: true},
	}

	expected := NewGameEngineState()
	expected.LastPlayedHand = nil
	expected.IsNewRound = true
	expected.PlayerPlacedInRound = false
	expected.Players = players{
		{Name: "Player 1", LastPlayed: false, IsPassed: false},
		{Name: "Player 2", LastPlayed: false, IsPassed: false},
		{Name: "Player 3", LastPlayed: false, IsPassed: false},
		{Name: "Player 4", LastPlayed: false, IsPassed: false},
	}

	engine := NewGameEngine()
	assert.Equal(t, expected, engine.setNewRound(state))
}

func TestProcessChangeNameRequest(t *testing.T) {
	req := changeNameRequest{PlayerName: "Player 5"}
	state := NewGameEngineState()
	state.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
		{Name: "Player 4"},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
		"003": {Player: state.Players[2]},
		"004": {Player: state.Players[3]},
	}

	expected := state
	expected.Players = players{
		{Name: "Player 1"},
		{Name: "Player 2"},
		{Name: "Player 3"},
		{Name: "Player 5"},
	}

	engine := NewGameEngine()
	assert.Equal(t, expected, engine.processChangeNameRequest(state, "004", req, senderMock{}))
}

// TestProcessTurnPlayRequest1 tests when a player places and wins the round
func TestProcessTurnPlayRequest1(t *testing.T) {
	state := NewGameEngineState()
	state.Players = players{
		{
			Name: "Player 1",
			Hand: []card{},
		},
		{
			Name: "Player 2",
			Hand: []card{},
		},
		{
			Name: "Player 3",
			Hand: []card{},
		},
		{
			Name: "Player 4",
			Hand: []card{},
		},
	}
	state.Connections = map[string]*PlayerContext{
		"001": {Player: state.Players[0]},
		"002": {Player: state.Players[1]},
		"003": {Player: state.Players[2]},
		"004": {Player: state.Players[3]},
	}

	// TODO: confirm with Harry if:
	//   a) it should always allow next non-placed to play even if passed, OR
	//   b) it should find next non-placed, non-passed to play (next non-placed if all passed)
}
