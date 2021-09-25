import { game, ConnectionState } from '~/store/game';
import { Card } from './models';
import {
  Message,
  JoinGameRequest,
  StartGameRequest,
  TurnPassRequest,
  TurnPlayRequest,
  ResetGameRequest,
} from './messages';

const wsUrl =
  window.location.protocol.replace('http', 'ws') +
  '//' +
  window.location.host +
  window.location.pathname +
  (window.location.pathname.endsWith('/') ? 'api' : '/api');

let socket: WebSocket | undefined = undefined;

export function joinGame({ name }: { name: string }): void {
  if (socket) socket.close();

  game.connState = ConnectionState.Connecting;

  socket = new WebSocket(wsUrl, 'json');
  socket.onmessage = onMessage;
  socket.onclose = onClose;
  socket.onerror = onError;
  socket.onopen = () => {
    game.connState = ConnectionState.Connected;
    const request: JoinGameRequest = { playerName: name };
    sendMessage({
      kind: 'JOIN_GAME',
      request,
    });
  };
}

export function requestStartGame(): void {
  const request: StartGameRequest = {};
  sendMessage({
    kind: 'START_GAME',
    request,
  });
}

export function requestResetGame(): void {
  const request: ResetGameRequest = {};
  sendMessage({
    kind: 'RESET_GAME',
    request,
  });
}

export function requestTurnPass(): void {
  const request: TurnPassRequest = {};
  sendMessage({
    kind: 'TURN_PASS',
    request,
  });
}

export function requestTurnPlay({ cards }: { cards: Card[] }): void {
  const request: TurnPlayRequest = { cards: cards.map((c) => c.globalRank) };
  sendMessage({
    kind: 'TURN_PLAY',
    request,
  });
}

function sendMessage({
  kind,
  request,
}: {
  kind: string;
  request:
    | JoinGameRequest
    | StartGameRequest
    | TurnPassRequest
    | TurnPlayRequest
    | ResetGameRequest;
}) {
  if (!socket || game.connState !== ConnectionState.Connected) return;
  const message: Message = {
    kind,
    data: btoa(JSON.stringify(request)),
  };
  socket.send(JSON.stringify(message));
}

// eslint-disable-next-line
const actions: Record<string, any> = {
  PLAYER_JOINED: game.playerJoined,
  PLAYER_DISCONNECTED: game.playerDisconnected,
  GAME_STARTED: game.gameStarted,
  GAME_PAUSED: game.gamePaused,
  GAME_RESUMED: game.gameResumed,
  GAME_RESET: game.gameReset,
  TURN_PASSED: game.turnPassed,
  ROUND_WON: game.roundWon,
  TURN_PLAYED: game.turnPlayed,
  PLAYER_PLACED: game.playerPlaced,
  GAME_WON: game.gameWon,
  GAME_STATE_REFRESH: game.gameStateRefresh,
  ERROR: game.actionError,
};

function onMessage(event: MessageEvent) {
  const message: Message | undefined = JSON.parse(event.data);
  if (!message) return;
  const parsed: unknown = JSON.parse(atob(message.data));
  actions[message.kind] && actions[message.kind]({ response: parsed });
}

function onClose() {
  game.selfDisconnected();
}

function onError(event: Event) {
  console.log('onError: ', event);
}
