import { ConnectionState } from '../stores/game'
import {
  type ChangeNameRequest,
  type ErrorResponse,
  type GamePausedResponse,
  type GameResetResponse,
  type GameResumedResponse,
  type GameStartedResponse,
  type GameStateRefreshResponse,
  type GameWonResponse,
  type JoinGameRequest,
  type Message,
  type NameChangedResponse,
  type PlayerDisconnectedResponse,
  type PlayerJoinedResponse,
  type PlayerPlacedResponse,
  type ResetGameRequest,
  type RoundWonResponse,
  type StartGameRequest,
  type TurnPassedResponse,
  type TurnPassRequest,
  type TurnPlayedResponse,
  type TurnPlayRequest,
} from './messages'
import { type Card } from './models'

export interface Store {
  connState: ConnectionState
  actionError: ({ response }: { response: ErrorResponse }) => void
  gamePaused: ({ _ }: { _: GamePausedResponse }) => void
  gameReset: ({ response }: { response: GameResetResponse }) => void
  gameResumed: ({ _ }: { _: GameResumedResponse }) => void
  gameStarted: ({ response }: { response: GameStartedResponse }) => void
  gameStateRefresh: ({ response }: { response: GameStateRefreshResponse }) => void
  gameWon: ({ response }: { response: GameWonResponse }) => void
  nameChanged: ({ response }: { response: NameChangedResponse }) => void
  playerDisconnected: ({ response }: { response: PlayerDisconnectedResponse }) => void
  playerJoined: ({ response }: { response: PlayerJoinedResponse }) => void
  playerPlaced: ({ response }: { response: PlayerPlacedResponse }) => void
  roundWon: ({ response }: { response: RoundWonResponse }) => void
  selfDisconnected: () => void
  turnPassed: ({ response }: { response: TurnPassedResponse }) => void
  turnPlayed: ({ response }: { response: TurnPlayedResponse }) => void
}

export interface Socket {
  joinGame: ({ name }: { name: string }) => void
  requestChangeName: ({ name }: { name: string }) => void
  requestResetGame: () => void
  requestStartGame: () => void
  requestTurnPass: () => void
  requestTurnPlay: ({ cards }: { cards: Card[] }) => void
}

export const init = (store: Store): Socket => {
  let socket: WebSocket | undefined = undefined

  const wsUrl =
    window.location.protocol.replace('http', 'ws') +
    '//' +
    window.location.host +
    window.location.pathname +
    (window.location.pathname.endsWith('/') ? 'api' : '/api')

  const joinGame = ({ name }: { name: string }): void => {
    if (socket) socket.close()

    store.connState = ConnectionState.Connecting

    socket = new WebSocket(wsUrl, 'json')
    socket.addEventListener('message', onMessage)
    socket.addEventListener('close', onClose)
    socket.addEventListener('error', onError)
    socket.addEventListener('open', () => {
      store.connState = ConnectionState.Connected
      const request: JoinGameRequest = { playerName: name }
      sendMessage({
        kind: 'JOIN_GAME',
        request,
      })
    })
  }

  const requestStartGame = (): void => {
    const request: StartGameRequest = {}
    sendMessage({
      kind: 'START_GAME',
      request,
    })
  }

  const requestResetGame = (): void => {
    const request: ResetGameRequest = {}
    sendMessage({
      kind: 'RESET_GAME',
      request,
    })
  }

  const requestTurnPass = (): void => {
    const request: TurnPassRequest = {}
    sendMessage({
      kind: 'TURN_PASS',
      request,
    })
  }

  const requestTurnPlay = ({ cards }: { cards: Card[] }): void => {
    const request: TurnPlayRequest = { cards: cards.map((c) => c.globalRank) }
    sendMessage({
      kind: 'TURN_PLAY',
      request,
    })
  }

  const requestChangeName = ({ name }: { name: string }): void => {
    const request: ChangeNameRequest = { name }
    sendMessage({
      kind: 'CHANGE_NAME',
      request,
    })
  }

  const sendMessage = ({
    kind,
    request,
  }: {
    kind: string
    request:
      | JoinGameRequest
      | StartGameRequest
      | TurnPassRequest
      | TurnPlayRequest
      | ResetGameRequest
      | ChangeNameRequest
  }) => {
    if (!socket || store.connState !== ConnectionState.Connected) return
    const message: Message = {
      kind,
      data: btoa(JSON.stringify(request)),
    }
    socket.send(JSON.stringify(message))
  }

  const onMessage = (event: MessageEvent) => {
    const message: Message | undefined = JSON.parse(event.data)
    if (!message) return

    const parsed: unknown = JSON.parse(atob(message.data))

    const actions: Record<string, Function> = {
      PLAYER_JOINED: store.playerJoined,
      PLAYER_DISCONNECTED: store.playerDisconnected,
      GAME_STARTED: store.gameStarted,
      GAME_PAUSED: store.gamePaused,
      GAME_RESUMED: store.gameResumed,
      GAME_RESET: store.gameReset,
      TURN_PASSED: store.turnPassed,
      ROUND_WON: store.roundWon,
      TURN_PLAYED: store.turnPlayed,
      NAME_CHANGED: store.nameChanged,
      PLAYER_PLACED: store.playerPlaced,
      GAME_WON: store.gameWon,
      GAME_STATE_REFRESH: store.gameStateRefresh,
      ERROR: store.actionError,
    }
    
    actions[message.kind] && actions[message.kind]({ response: parsed })
  }

  const onClose = () => {
    store.selfDisconnected()
  }

  const onError = (event: Event) => {
    console.log('onError: ', event)
  }

  return {
    joinGame,
    requestChangeName,
    requestResetGame,
    requestStartGame,
    requestTurnPass,
    requestTurnPlay,
  }
}
