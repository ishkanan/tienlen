import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Score from '../views/Score.vue'
import { useGameStore } from '../stores/game'
import App from '../App.vue'

import './Score.stories.css'

const meta = {
  title: 'Score',
  component: Score,
  decorators: [() => ({
    template: '<div class="scoreArea"><story/></div>'
  })],
  args: { },
} satisfies Meta<typeof Score>

export default meta
type Story = StoryObj<typeof meta>

const app = createApp(App)

app.use(createPinia())

const gameStore = useGameStore()

const reset = () => {
  gameStore.self = {
    name: 'Mario',
    position: 1,
    cardsLeft: 13,
    isPassed: false,
    isTurn: true,
    wonLastGame: false,
    connected: true,
    score: 5,
    lastPlayed: false,
  }

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
}

export const FourPlayers: Story = {
  play: async () => {
    reset()
    gameStore.winPlaces = []
  }
}

export const FourPlayersEndGame: Story = {
  play: async () => {
    reset()
    gameStore.winPlaces = gameStore.opponents.concat(gameStore.self!)
  }
}
