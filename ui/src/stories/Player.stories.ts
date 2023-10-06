import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Player from '../views/Player.vue'
import { useGameStore } from '../stores/game'
import App from '../App.vue'
import { generateDeck } from './Hand.stories'

import './Player.stories.css'
import { GameState } from '@/lib/messages'

const meta = {
  title: 'Player',
  component: Player,
  decorators: [() => ({
    template: '<div class="playerArea"><story/></div>'
  })],
  args: { },
} satisfies Meta<typeof Player>

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
    cardsLeft: 13,
    isPassed: false,
    isTurn: false,
    wonLastGame: false,
    connected: true,
    score: 0,
    lastPlayed: false,
  }

  gameStore.gameState = GameState.InLobby

  gameStore.selfHand = []

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

  gameStore.winPlaces = []
}

export const GameCanStart: Story = {
  loaders: [
    async () => ({
      work: function () {
        reset()
        gameStore.selfHand = deck.slice(0, 14)
      }()
    })
  ]
}
