import { useGameStore, ConnectionState } from '../stores/game'
import { type Card } from './models'
import {
  type Message,
  type JoinGameRequest,
  type StartGameRequest,
  type TurnPassRequest,
  type TurnPlayRequest,
  type ResetGameRequest,
  type ChangeNameRequest,
} from './messages'

export interface Socket {
  joinGame: ({ name }: { name: string }) => void
  requestChangeName: ({ name }: { name: string }) => void
  requestResetGame: () => void
  requestStartGame: () => void
  requestTurnPass: () => void
  requestTurnPlay: ({ cards }: { cards: Card[] }) => void
}

export function init(): Socket {
  const wsUrl =
    window.location.protocol.replace('http', 'ws') +
    '//' +
    window.location.host +
    window.location.pathname +
    (window.location.pathname.endsWith('/') ? 'api' : '/api')

  let gameState = useGameStore()
  let socket: WebSocket | undefined = undefined

  const joinGame = ({ name }: { name: string }): void => {
    if (socket) socket.close()

    gameState.connState = ConnectionState.Connecting

    socket = new WebSocket(wsUrl, 'json')
    socket.onmessage = onMessage
    socket.onclose = onClose
    socket.onerror = onError
    socket.onopen = () => {
      gameState.connState = ConnectionState.Connected
      const request: JoinGameRequest = { playerName: name }
      sendMessage({
        kind: 'JOIN_GAME',
        request,
      })
    }
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
    if (!socket || gameState.connState !== ConnectionState.Connected) return
    const message: Message = {
      kind,
      data: btoa(JSON.stringify(request)),
    }
    socket.send(JSON.stringify(message))
  }

  // eslint-disable-next-line
  const actions: Record<string, any> = {
    PLAYER_JOINED: gameState.playerJoined,
    PLAYER_DISCONNECTED: gameState.playerDisconnected,
    GAME_STARTED: gameState.gameStarted,
    GAME_PAUSED: gameState.gamePaused,
    GAME_RESUMED: gameState.gameResumed,
    GAME_RESET: gameState.gameReset,
    TURN_PASSED: gameState.turnPassed,
    ROUND_WON: gameState.roundWon,
    TURN_PLAYED: gameState.turnPlayed,
    NAME_CHANGED: gameState.nameChanged,
    PLAYER_PLACED: gameState.playerPlaced,
    GAME_WON: gameState.gameWon,
    GAME_STATE_REFRESH: gameState.gameStateRefresh,
    ERROR: gameState.actionError,
  }

  const onMessage = (event: MessageEvent) => {
    const message: Message | undefined = JSON.parse(event.data)
    if (!message) return
    const parsed: unknown = JSON.parse(atob(message.data))
    actions[message.kind] && actions[message.kind]({ response: parsed })
  }

  const onClose = () => {
    gameState.selfDisconnected()
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
