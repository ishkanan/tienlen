import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Discard from '../views/Discard.vue'
import { useGameStore } from '../stores/game'
import { GameState } from '../lib/messages'
import App from '../App.vue'
import { generateDeck } from './Hand.stories'

import './Discard.stories.css'

const meta = {
  title: 'Discard',
  component: Discard,
  decorators: [() => ({
    template: '<div class="discardArea"><story/></div>'
  })],
  args: { },
} satisfies Meta<typeof Discard>

export default meta
type Story = StoryObj<typeof meta>

const app = createApp(App)

app.use(createPinia())

const deck = generateDeck()

const gameStore = useGameStore()

const reset = () => {
  gameStore.self = {
    name: 'Mario',
    position: 1,
    cardsLeft: 0,
    isPassed: false,
    isTurn: false,
    wonLastGame: false,
    connected: true,
    score: 0,
    lastPlayed: false,
  }  
}

export const PausedGame: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.Paused
  }
}

export const FirstGameLobbyNoOpponents: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.InLobby
    gameStore.opponents = []
    gameStore.self!.wonLastGame = false
  }
}

export const FirstGameLobbySomeOpponents: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.InLobby
    gameStore.opponents = [
      {
        name: 'Luigi',
        position: 2,
        cardsLeft: 13,
        isPassed: false,
        isTurn: false,
        wonLastGame: false,
        connected: true,
        score: 3,
        lastPlayed: false,
      }
    ]
    gameStore.self!.wonLastGame = false
  }
}

export const FirstGameLobbyAllOpponents: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.InLobby
    gameStore.opponents = [
      {
        name: 'Luigi',
        position: 2,
        cardsLeft: 13,
        isPassed: false,
        isTurn: false,
        wonLastGame: false,
        connected: true,
        score: 3,
        lastPlayed: false,
      },
      {
        name: 'Peach',
        position: 3,
        cardsLeft: 13,
        isPassed: false,
        isTurn: false,
        wonLastGame: false,
        connected: true,
        score: 8,
        lastPlayed: false,
      },
      {
        name: 'Toad',
        position: 4,
        cardsLeft: 13,
        isPassed: false,
        isTurn: false,
        wonLastGame: false,
        connected: true,
        score: 12,
        lastPlayed: false,
      },
    ]
    gameStore.self!.wonLastGame = false
  }
}

export const NextGameLobbyWaitingToStart: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.InLobby
    gameStore.self!.wonLastGame = true
    gameStore.lastPlayed = []
  }
}

export const InProgressNoCardsPlayed: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.Running
    gameStore.lastPlayed = []
  }
}

export const InProgressSomeCardsPlayed: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.Running
    gameStore.lastPlayed = deck.slice(0, 5)
    gameStore.self!.lastPlayed = true
  }
}

export const InProgressManyCardsPlayed: Story = {
  play: async () => {
    reset()
    gameStore.gameState = GameState.Running
    gameStore.lastPlayed = deck.slice(0, 13)
    gameStore.self!.lastPlayed = true
  }
}
