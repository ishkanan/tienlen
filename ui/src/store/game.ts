import { VuexModule, Module, Action } from 'vuex-class-modules';
import {
  ErrorKind,
  ErrorResponse,
  GamePausedResponse,
  GameResumedResponse,
  GameStartedResponse,
  GameState,
  GameStateRefreshResponse,
  GameWonResponse,
  PlayerDisconnectedResponse,
  PlayerJoinedResponse,
  PlayerPlacedResponse,
  RoundWonResponse,
  TurnPassedResponse,
  TurnPlayedResponse,
} from '~/lib/messages';
import { Card, GameEvent, GameEventKind, Player } from '~/lib/models';
import { ordinalise } from '~/lib/utils';
import store from '~/store';

export enum ConnectionState {
  NotConnected = 0,
  Connecting = 1,
  Connected = 2
}

@Module({ generateMutationSetters: true })
class Game extends VuexModule {
  connState: ConnectionState = ConnectionState.NotConnected;
  name = '';
  events: GameEvent[] = [];
  opponents: Player[] = [];
  self: Player | undefined = undefined;
  selfHand: Card[] = [];
  gameState: GameState = GameState.InLobby;
  lastPlayed: Card[] = [];
  firstRound = true;
  newRound = true;
  winPlaces: Player[] = [];

  private errorMap = {
    [ErrorKind.LobbyNotReady]: 'Game is not ready to start yet.',
    [ErrorKind.NotAuthorised]: 'You cannot perform that action now.',
    [ErrorKind.OutOfTurn]: 'It is not your turn.',
    [ErrorKind.MustPlay]: 'You must play one or more cards.',
    [ErrorKind.InvalidCards]: 'Game did not recognise your cards.',
    [ErrorKind.InvalidPattern]: 'Those cards cannot be played.',
    [ErrorKind.CardsNotBetter]: 'Cards do not beat last played.',
    [ErrorKind.MustPlayLowest]: 'Must play lowest card.',
    [ErrorKind.NameTaken]: 'That name is already taken.',
    [ErrorKind.GameFull]: 'The game is full.',
  };

  get isInLobby(): boolean {
    return this.gameState === GameState.InLobby;
  }

  get isInProgress(): boolean {
    return this.gameState !== GameState.InLobby;
  }

  get isPaused(): boolean {
    return this.gameState === GameState.Paused;
  }

  @Action
  playerJoined({ response }: { response: PlayerJoinedResponse }) {
    this.events.push({
      kind: GameEventKind.Info,
      message: `${response.player.name} has joined the game.`,
    });
    this.connState = ConnectionState.Connected;
  }

  @Action
  selfDisconnected() {
    this.events.push({
      kind: GameEventKind.Error,
      message: 'You were disconnected from the game.',
    });
    this.connState = ConnectionState.NotConnected;
    this.firstRound = true;
    this.gameState = GameState.InLobby;
    this.lastPlayed = [];
    this.newRound = true;
    this.opponents = [];
    this.self = undefined;
    this.selfHand = [];
    this.winPlaces = [];
  }

  @Action
  playerDisconnected({ response }: { response: PlayerDisconnectedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has left the game.`,
    });
  }

  @Action
  gameStarted({ response }: { response: GameStartedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has started the game.`,
    });
  }

  @Action
  gamePaused({ _ }: { _: GamePausedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: 'Game paused, will resume when all players re-connect.',
    });
  }

  @Action
  gameResumed({ _ }: { _: GameResumedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: 'All players have re-connected, game resumed.',
    });
  }

  @Action
  turnPassed({ response }: { response: TurnPassedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has passed their turn.`,
    });
  }

  @Action
  playerPlaced({ response }: { response: PlayerPlacedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has no more cards and got ${ordinalise(response.place)} place.`,
    });
  }

  @Action
  roundWon({ response }: { response: RoundWonResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has won the round.`,
    });
  }

  @Action
  turnPlayed({ response }: { response: TurnPlayedResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has played their turn.`,
    });
  }

  @Action
  gameWon({ response }: { response: GameWonResponse }) {
    this.showNotification({
      kind: GameEventKind.Info,
      message: `${response.player.name} has won the game.`,
    });
  }

  @Action
  gameStateRefresh({ response }: { response: GameStateRefreshResponse }) {
    this.name = this.self ? this.self.name : this.name;
    this.firstRound = response.firstRound;
    this.gameState = response.gameState;
    this.lastPlayed = response.lastPlayed ?? [];
    this.newRound = response.newRound;
    this.opponents = response.opponents;
    this.self = response.self;
    this.selfHand = response.selfHand;
    this.winPlaces = response.winPlaces;
  }

  @Action
  actionError({ response }: { response: ErrorResponse }) {
    this.showNotification({
      kind: GameEventKind.Error,
      message: this.errorMap[response.kind],
    });
  }

  @Action
  showNotification({ kind, message }: { kind: GameEventKind, message: string }) {
    this.events.push({ kind, message });
  }
}

export const game = new Game({ store, name: 'game' });
