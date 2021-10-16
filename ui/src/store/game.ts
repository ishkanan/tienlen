import { orderBy } from 'lodash-es';
import { VuexModule, Module, Action } from 'vuex-class-modules';
import {
  ErrorKind,
  ErrorResponse,
  GamePausedResponse,
  GameResetResponse,
  GameResumedResponse,
  GameStartedResponse,
  GameState,
  GameStateRefreshResponse,
  GameWonResponse,
  NameChangedResponse,
  PlayerDisconnectedResponse,
  PlayerJoinedResponse,
  PlayerPlacedResponse,
  RoundWonResponse,
  TurnPassedResponse,
  TurnPlayedResponse,
} from '~/lib/messages';
import { Card, Event, EventRune, EventSeverity, Player } from '~/lib/models';
import { ordinalise } from '~/lib/utils';
import store from '~/store';

export enum ConnectionState {
  NotConnected = 0,
  Connecting = 1,
  Connected = 2,
}

@Module({ generateMutationSetters: true })
class Game extends VuexModule {
  connState: ConnectionState = ConnectionState.NotConnected;
  name = '';
  events: Event[] = [];
  opponents: Player[] = [];
  self: Player | undefined = undefined;
  selfHand: Card[] = [];
  gameState: GameState = GameState.InLobby;
  lastPlayed: Card[] = [];
  firstRound = true;
  newRound = true;
  winPlaces: Player[] = [];

  private errorMap = {
    [ErrorKind.LobbyNotReady]: {
      message: 'Game is not yet ready to start.',
      toast: false,
    },
    [ErrorKind.NotAuthorised]: {
      message: 'You cannot perform that action now.',
      toast: false,
    },
    [ErrorKind.OutOfTurn]: {
      message: 'It is not your turn.',
      toast: false,
    },
    [ErrorKind.MustPlay]: {
      message: 'You must play one or more cards.',
      toast: false,
    },
    [ErrorKind.InvalidCards]: {
      message: 'Unrecognised cards, did you select any?',
      toast: false,
    },
    [ErrorKind.InvalidPattern]: {
      message: 'Cards cannot be played.',
      toast: false,
    },
    [ErrorKind.CardsNotBetter]: {
      message: 'Cards do not beat the last played hand.',
      toast: false,
    },
    [ErrorKind.MustPlayLowest]: {
      message: 'You must play your lowest card.',
      toast: false,
    },
    [ErrorKind.InvalidName]: {
      message: 'That name is not valid.',
      toast: (context: this) => context.connState === ConnectionState.NotConnected,
    },
    [ErrorKind.GameFull]: {
      message: 'The game is full.',
      toast: true,
    },
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

  get isFirstGame(): boolean {
    const scores = this.opponents.reduce((memo, p) => memo + p.score, this.self?.score ?? 0);
    return scores === 0;
  }

  @Action
  playerJoined({ response }: { response: PlayerJoinedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has joined the game.` }],
    });
    this.connState = ConnectionState.Connected;
  }

  @Action
  selfDisconnected() {
    this.pushEvent({
      severity: EventSeverity.Error,
      runes: [{ message: 'You were disconnected from the game.' }],
      toast: true,
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
    this.pushEvent({
      severity: this.isInLobby ? EventSeverity.Info : EventSeverity.Error,
      runes: [{ message: `"${response.player.name}" has left the game.` }],
    });
  }

  @Action
  gameStarted({ response }: { response: GameStartedResponse }) {
    if (!this.isFirstGame) this.events = [];
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has started the game.` }],
    });
  }

  @Action
  gamePaused({ _ }: { _: GamePausedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Warning,
      runes: [{ message: 'Game is paused and will resume when all players re-connect.' }],
    });
  }

  @Action
  gameResumed({ _ }: { _: GameResumedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: 'All players have re-connected, game has resumed.' }],
    });
  }

  @Action
  turnPassed({ response }: { response: TurnPassedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has passed their turn.` }],
    });
  }

  @Action
  playerPlaced({ response }: { response: PlayerPlacedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Success,
      runes: [
        {
          message: `"${response.player.name}" has no more cards and placed ${ordinalise(
            response.place,
          )}.`,
        },
      ],
    });
  }

  @Action
  roundWon({ response }: { response: RoundWonResponse }) {
    this.pushEvent({
      severity: EventSeverity.Success,
      runes: [{ message: `"${response.player.name}" has won the round.` }],
    });
  }

  @Action
  turnPlayed({ response }: { response: TurnPlayedResponse }) {
    const cards = orderBy(response.cards, ['globalRank'], ['desc']).map((c) => ({ card: c }));
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [
        {
          message: `"${response.player.name}" played `,
        },
        ...cards,
      ],
    });
  }

  @Action
  nameChanged({ response }: { response: NameChangedResponse }) {
    this.pushEvent({
      severity: EventSeverity.Info,
      runes: [
        { message: `"${response.oldPlayer.name}" is now known as "${response.newPlayer.name}"` },
      ],
    });
  }

  @Action
  gameWon({ response }: { response: GameWonResponse }) {
    this.pushEvent({
      severity: EventSeverity.Success,
      runes: [{ message: `"${response.player.name}" has won the game.` }],
    });
  }

  @Action
  gameReset({ response }: { response: GameResetResponse }) {
    this.pushEvent({
      severity: EventSeverity.Warning,
      runes: [{ message: `"${response.player.name}" has reset the game. Dick move?` }],
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
    this.pushEvent({
      severity: EventSeverity.Error,
      runes: [{ message: this.errorMap[response.kind].message }],
      toast: (function (context) {
        const t = context.errorMap[response.kind].toast;
        if (typeof t === 'boolean') return t;
        return t(context);
      })(this),
    });
  }

  @Action
  pushEvent({
    severity,
    runes,
    toast,
  }: {
    severity: EventSeverity;
    runes: EventRune[];
    toast?: boolean;
  }) {
    this.events.push({
      severity,
      runes,
      timestamp: new Date(),
      toast: toast !== undefined ? toast : false,
    });
  }
}

export const game = new Game({ store, name: 'game' });
