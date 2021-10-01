package game

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/ishkanan/tienlen/api/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:  func(r *http.Request) bool { return true },
	Subprotocols: []string{"json"},
}

// ISender defines how the game pumps events to the underlying sockets
type ISender interface {
	SendResponseToAllPlayers(state GameEngineState, response interface{}) map[*PlayerContext]error
	SendResponseOnConnection(state GameEngineState, connID string, response interface{}) error
	SendStateToAllPlayers(state GameEngineState) map[*PlayerContext]error
}

// IMessageSource defines how the ingress socket events are pumped to the game
type IMessageSource interface {
	IsAcceptingConnections(GameEngineState) bool
	ConnectionStateChanged(GameEngineState, uuid.UUID, IMessageSink, connState, ISender) GameEngineState
	ProcessRequest(GameEngineState, uuid.UUID, interface{}, reflect.Type, ISender) GameEngineState
}

// IMessageSink defines how the game interfaces with a single underlying socket
type IMessageSink interface {
	Send(interface{}) error
	Close() error
}

// MessageSink provides an implementation of the IMessageSink interface
type MessageSink struct {
	ConnID     uuid.UUID
	Connection *websocket.Conn
}

// Send attempts to send a response-type message through the underlying connection
func (m *MessageSink) Send(response interface{}) error {
	responseType := reflect.TypeOf(response)
	for t, ident := range responseMap() {
		if t == responseType {
			responseBytes, err := json.Marshal(response)
			if err != nil {
				utils.LogDebug("Send:: Marshal error for %s - %v", m.ConnID.String(), err)
				return err
			}
			return m.Connection.WriteJSON(Message{
				Kind: ident,
				Data: base64.StdEncoding.EncodeToString(responseBytes),
			})
		}
	}
	return fmt.Errorf("unrecognised response type (%s)", responseType.Name())
}

// Close closes the underlying connection
func (m *MessageSink) Close() error {
	return m.Connection.Close()
}

// builds a request (ingress message) object based on the identifier
func buildRequest(ident string, data []byte) (interface{}, error) {
	switch ident {
	case "JOIN_GAME":
		request := joinGameRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	case "START_GAME":
		request := startGameRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	case "RESET_GAME":
		request := resetGameRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	case "TURN_PASS":
		request := turnPassRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	case "TURN_PLAY":
		request := turnPlayRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	case "CHANGE_NAME":
		request := changeNameRequest{}
		err := json.Unmarshal(data, &request)
		return request, err
	default:
		return nil, fmt.Errorf("unrecognised request type (%s)", ident)
	}
}

// maps response (egress messages) Golang types to identifiers
func responseMap() map[reflect.Type]string {
	return map[reflect.Type]string{
		reflect.TypeOf(playerJoinedResponse{}):       "PLAYER_JOINED",
		reflect.TypeOf(playerDisconnectedResponse{}): "PLAYER_DISCONNECTED",
		reflect.TypeOf(gameStartedResponse{}):        "GAME_STARTED",
		reflect.TypeOf(gamePausedResponse{}):         "GAME_PAUSED",
		reflect.TypeOf(gameResumedResponse{}):        "GAME_RESUMED",
		reflect.TypeOf(gameResetResponse{}):          "GAME_RESET",
		reflect.TypeOf(turnPassedResponse{}):         "TURN_PASSED",
		reflect.TypeOf(roundWonResponse{}):           "ROUND_WON",
		reflect.TypeOf(turnPlayedResponse{}):         "TURN_PLAYED",
		reflect.TypeOf(nameChangedResponse{}):        "NAME_CHANGED",
		reflect.TypeOf(playerPlacedResponse{}):       "PLAYER_PLACED",
		reflect.TypeOf(gameWonResponse{}):            "GAME_WON",
		reflect.TypeOf(gameStateRefreshResponse{}):   "GAME_STATE_REFRESH",
		reflect.TypeOf(errorResponse{}):              "ERROR",
	}
}

// ConnectionHandler provides incoming message and ping "pump" logic for a given connection
func ConnectionHandler(game IMessageSource, initState GameEngineState, pump ISender) func(w http.ResponseWriter, r *http.Request) {
	state := initState
	mutex := &sync.Mutex{}

	isAcceptingConnections := func() bool {
		mutex.Lock()
		defer mutex.Unlock()
		return game.IsAcceptingConnections(state)
	}

	connectionStateChanged := func(connID uuid.UUID, sink IMessageSink, connState connState) GameEngineState {
		mutex.Lock()
		defer mutex.Unlock()
		return game.ConnectionStateChanged(state, connID, sink, connState, pump)
	}

	processRequest := func(connID uuid.UUID, request interface{}, requestType reflect.Type) GameEngineState {
		mutex.Lock()
		defer mutex.Unlock()
		return game.ProcessRequest(state, connID, request, requestType, pump)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		utils.LogDebug("ConnectionHandler:: New HTTP connection from %s", r.RemoteAddr)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			utils.LogDebug("ConnectionHandler:: socket upgrade error for %s - %v", r.RemoteAddr, err)
			return
		}
		defer conn.Close()

		connID := uuid.New()
		utils.LogDebug("ConnectionHandler:: %s is assigned connID %s", r.RemoteAddr, connID.String())
		sink := &MessageSink{ConnID: connID, Connection: conn}

		if !isAcceptingConnections() {
			_ = sink.Send(errorResponse{Kind: errKindGameFull})
			utils.LogDebug("ConnectionHandler:: %s tried to join, but game is full", connID.String())
			return
		}

		state = connectionStateChanged(connID, sink, connStateNew)

		var lastPingError error
		conn.SetPongHandler(func(appData string) error {
			go func() {
				// sends one final ping after connection death, which is fine
				time.Sleep(5 * time.Second)
				lastPingError = conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(5*time.Second))
			}()
			if lastPingError != nil {
				utils.LogDebug("ConnectionHandler:: ping write error for %s - %v", connID.String(), err)
			}
			return lastPingError
		})

		err = conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(5*time.Second))
		if err != nil {
			utils.LogDebug("ConnectionHandler:: initial ping write error for %s - %v", connID.String(), err)
			state = connectionStateChanged(connID, sink, connStateDead)
			return
		}

		var message Message
		for {
			if conn.ReadJSON(&message); err != nil {
				utils.LogDebug("ConnectionHandler:: read error for %s - %v", connID.String(), err)
				state = connectionStateChanged(connID, sink, connStateDead)
				return
			}

			requestBytes, err := base64.StdEncoding.DecodeString(message.Data)
			if err != nil {
				utils.LogDebug("ConnectionHandler:: decode error for %s - %v", connID.String(), err)
				continue
			}

			request, err := buildRequest(message.Kind, requestBytes)
			if err != nil {
				utils.LogDebug("ConnectionHandler:: marshal error for %s - %v", connID.String(), err)
				continue
			}

			state = processRequest(connID, request, reflect.TypeOf(request))
		}
	}
}
