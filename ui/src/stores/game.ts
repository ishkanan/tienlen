import { defineStore } from 'pinia'
import { orderBy } from 'lodash-es'
import {
  ErrorKind,
  type ErrorResponse,
  type GamePausedResponse,
  type GameResetResponse,
  type GameResumedResponse,
  type GameStartedResponse,
  GameState,
  type GameStateRefreshResponse,
  type GameWonResponse,
  type NameChangedResponse,
  type PlayerDisconnectedResponse,
  type PlayerJoinedResponse,
  type PlayerPlacedResponse,
  type RoundWonResponse,
  type TurnPassedResponse,
  type TurnPlayedResponse,
} from '../lib/messages'
import {
  type Card,
  type Event,
  type EventRune,
  EventSeverity,
  type Player
} from '../lib/models'
import { ordinalise } from '../lib/utils'

export enum ConnectionState {
  NotConnected = 0,
  Connecting = 1,
  Connected = 2,
}

const errorMap = {
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
    toast: (connState: ConnectionState) => connState === ConnectionState.NotConnected,
  },
  [ErrorKind.GameFull]: {
    message: 'The game is full.',
    toast: true,
  },
}

export const useGameStore = defineStore('game', {
  state: () => ({
    connState: ConnectionState.NotConnected,
    events: [] as Event[],
    firstRound: true,
    gameState: GameState.InLobby,
    lastPlayed: [] as Card[],
    name: '',
    newRound: true,
    opponents: [] as Player[],
    self: undefined as Player | undefined,
    selfHand: [] as Card[],
    winPlaces: [] as Player[],
  }),

  getters: {
    isInLobby: ({ gameState }) => gameState === GameState.InLobby,
    isInProgress: ({ gameState }) => gameState !== GameState.InLobby,
    isPaused: ({ gameState }) => gameState === GameState.Paused,
    isFirstGame: ({ opponents, self }) => {
      const scores = opponents.reduce((memo, p) => memo + p.score, self?.score ?? 0)
      return scores === 0
    }
  },

  actions: {
    actionError({ response }: { response: ErrorResponse }) {
      this.pushEvent({
        severity: EventSeverity.Error,
        runes: [{ message: errorMap[response.kind].message }],
        toast: (() => {
          const t = errorMap[response.kind].toast
          if (typeof t === 'boolean') return t
          return t(this.connState);
        })(),
      })
    },
    gamePaused({ _ }: { _: GamePausedResponse }) {
      this.pushEvent({
        severity: EventSeverity.Warning,
        runes: [{ message: 'Game is paused and will resume when all players re-connect.' }],
      })
    },
    gameReset({ response }: { response: GameResetResponse }) {
      this.pushEvent({
        severity: EventSeverity.Warning,
        runes: [{ message: `"${response.player.name}" has reset the game. Dick move?` }],
      })
    },
    gameResumed({ _ }: { _: GameResumedResponse }) {
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [{ message: 'All players have re-connected, game has resumed.' }],
      })
    },
    gameStarted({ response }: { response: GameStartedResponse }) {
      if (!this.isFirstGame) this.events = []
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [{ message: `"${response.player.name}" has started the game.` }],
      })
    },
    gameStateRefresh({ response }: { response: GameStateRefreshResponse }) {
      this.name = self?.name ?? this.name
      this.firstRound = response.firstRound
      this.gameState = response.gameState
      this.lastPlayed = response.lastPlayed ?? []
      this.newRound = response.newRound
      this.opponents = response.opponents
      this.self = response.self
      this.selfHand = response.selfHand
      this.winPlaces = response.winPlaces
    },
    gameWon({ response }: { response: GameWonResponse }) {
      this.pushEvent({
        severity: EventSeverity.Success,
        runes: [{ message: `"${response.player.name}" has won the game.` }],
      })
    },
    nameChanged({ response }: { response: NameChangedResponse }) {
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [
          { message: `"${response.oldPlayer.name}" is now known as "${response.newPlayer.name}"` },
        ],
      })
    },
    playerDisconnected({ response }: { response: PlayerDisconnectedResponse }) {
      this.pushEvent({
        severity: this.isInLobby ? EventSeverity.Info : EventSeverity.Error,
        runes: [{ message: `"${response.player.name}" has left the game.` }],
      })
    },
    playerJoined({ response }: { response: PlayerJoinedResponse }) {
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [{ message: `"${response.player.name}" has joined the game.` }],
      })
      this.connState = ConnectionState.Connected
    },
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
      })
    },
    pushEvent({
      severity,
      runes,
      toast,
    }: {
      severity: EventSeverity
      runes: EventRune[]
      toast?: boolean
    }) {
      this.events.push({
        severity,
        runes,
        timestamp: new Date(),
        toast: toast !== undefined ? toast : false,
      })
    },
    roundWon({ response }: { response: RoundWonResponse }) {
      this.pushEvent({
        severity: EventSeverity.Success,
        runes: [{ message: `"${response.player.name}" has won the round.` }],
      })
    },
    selfDisconnected() {
      this.pushEvent({
        severity: EventSeverity.Error,
        runes: [{ message: 'You were disconnected from the game.' }],
        toast: true,
      })
      this.connState = ConnectionState.NotConnected
      this.firstRound = true
      this.gameState = GameState.InLobby
      this.lastPlayed = []
      this.newRound = true
      this.opponents = []
      this.self = undefined
      this.selfHand = []
      this.winPlaces = []
    },
    turnPassed({ response }: { response: TurnPassedResponse }) {
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [{ message: `"${response.player.name}" has passed their turn.` }],
      })
    },
    turnPlayed({ response }: { response: TurnPlayedResponse }) {
      const cards = orderBy(response.cards, ['globalRank'], ['desc']).map((c) => ({ card: c }))
      this.pushEvent({
        severity: EventSeverity.Info,
        runes: [
          {
            message: `"${response.player.name}" played `,
          },
          ...cards,
        ],
      })
    }
  }
})
