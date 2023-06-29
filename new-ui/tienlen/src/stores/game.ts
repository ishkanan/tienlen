import { ref, computed } from 'vue'
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

export const useGameStore = defineStore('game', () => {
  const connState = ref(ConnectionState.NotConnected)
  const name = ref('')
  const events = ref([] as Event[])
  const opponents = ref([] as Player[])
  const self = ref(undefined as Player | undefined)
  const selfHand = ref([] as Card[])
  const gameState = ref(GameState.InLobby)
  const lastPlayed = ref([] as Card[])
  const firstRound = ref(true)
  const newRound = ref(true)
  const winPlaces = ref([] as Player[])

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
      toast: () => connState.value === ConnectionState.NotConnected,
    },
    [ErrorKind.GameFull]: {
      message: 'The game is full.',
      toast: true,
    },
  }

  const isInLobby = computed(() => {
    return gameState.value === GameState.InLobby
  })

  const isInProgress = computed(() => {
    return gameState.value !== GameState.InLobby
  })

  const isPaused = computed(() => {
    return gameState.value === GameState.Paused
  })

  const isFirstGame = computed(() => {
    const scores = opponents.value.reduce((memo, p) => memo + p.score, self?.value?.score ?? 0)
    return scores === 0
  })

  function playerJoined({ response }: { response: PlayerJoinedResponse }) {
    pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has joined the game.` }],
    })
    connState.value = ConnectionState.Connected
  }

  function selfDisconnected() {
    pushEvent({
      severity: EventSeverity.Error,
      runes: [{ message: 'You were disconnected from the game.' }],
      toast: true,
    })
    connState.value = ConnectionState.NotConnected
    firstRound.value = true
    gameState.value = GameState.InLobby
    lastPlayed.value = []
    newRound.value = true
    opponents.value = []
    self.value = undefined
    selfHand.value = []
    winPlaces.value = []
  }

  function playerDisconnected({ response }: { response: PlayerDisconnectedResponse }) {
    pushEvent({
      severity: isInLobby.value ? EventSeverity.Info : EventSeverity.Error,
      runes: [{ message: `"${response.player.name}" has left the game.` }],
    })
  }

  function gameStarted({ response }: { response: GameStartedResponse }) {
    if (!isFirstGame.value) events.value = []
    pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has started the game.` }],
    })
  }

  function gamePaused({ _ }: { _: GamePausedResponse }) {
    pushEvent({
      severity: EventSeverity.Warning,
      runes: [{ message: 'Game is paused and will resume when all players re-connect.' }],
    })
  }

  function gameResumed({ _ }: { _: GameResumedResponse }) {
    pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: 'All players have re-connected, game has resumed.' }],
    })
  }

  function turnPassed({ response }: { response: TurnPassedResponse }) {
    pushEvent({
      severity: EventSeverity.Info,
      runes: [{ message: `"${response.player.name}" has passed their turn.` }],
    })
  }

  function playerPlaced({ response }: { response: PlayerPlacedResponse }) {
    pushEvent({
      severity: EventSeverity.Success,
      runes: [
        {
          message: `"${response.player.name}" has no more cards and placed ${ordinalise(
            response.place,
          )}.`,
        },
      ],
    })
  }

  function roundWon({ response }: { response: RoundWonResponse }) {
    pushEvent({
      severity: EventSeverity.Success,
      runes: [{ message: `"${response.player.name}" has won the round.` }],
    })
  }

  function turnPlayed({ response }: { response: TurnPlayedResponse }) {
    const cards = orderBy(response.cards, ['globalRank'], ['desc']).map((c) => ({ card: c }))
    pushEvent({
      severity: EventSeverity.Info,
      runes: [
        {
          message: `"${response.player.name}" played `,
        },
        ...cards,
      ],
    })
  }

  function nameChanged({ response }: { response: NameChangedResponse }) {
    pushEvent({
      severity: EventSeverity.Info,
      runes: [
        { message: `"${response.oldPlayer.name}" is now known as "${response.newPlayer.name}"` },
      ],
    })
  }

  function gameWon({ response }: { response: GameWonResponse }) {
    pushEvent({
      severity: EventSeverity.Success,
      runes: [{ message: `"${response.player.name}" has won the game.` }],
    })
  }

  function gameReset({ response }: { response: GameResetResponse }) {
    pushEvent({
      severity: EventSeverity.Warning,
      runes: [{ message: `"${response.player.name}" has reset the game. Dick move?` }],
    })
  }

  function gameStateRefresh({ response }: { response: GameStateRefreshResponse }) {
    name.value = self?.value?.name ?? name.value
    firstRound.value = response.firstRound
    gameState.value = response.gameState
    lastPlayed.value = response.lastPlayed ?? []
    newRound.value = response.newRound
    opponents.value = response.opponents
    self.value = response.self
    selfHand.value = response.selfHand
    winPlaces.value = response.winPlaces
  }

  function actionError({ response }: { response: ErrorResponse }) {
    pushEvent({
      severity: EventSeverity.Error,
      runes: [{ message: errorMap[response.kind].message }],
      toast: (() => {
        const t = errorMap[response.kind].toast;
        if (typeof t === 'boolean') return t;
        return t();
      })(),
    });
  }

  function pushEvent({
    severity,
    runes,
    toast,
  }: {
    severity: EventSeverity
    runes: EventRune[]
    toast?: boolean
  }) {
    events.value.push({
      severity,
      runes,
      timestamp: new Date(),
      toast: toast !== undefined ? toast : false,
    })
  }

  return {
    connState,
    events,
    isInLobby,
    isInProgress,
    isPaused,
    lastPlayed,
    name,
    opponents,
    self,
    winPlaces,
    actionError,
    gamePaused,
    gameReset,
    gameResumed,
    gameStarted,
    gameStateRefresh,
    gameWon,
    nameChanged,
    playerDisconnected,
    playerJoined,
    playerPlaced,
    roundWon,
    selfDisconnected,
    turnPassed,
    turnPlayed,
  }
})
