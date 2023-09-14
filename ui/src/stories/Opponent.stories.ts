import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Opponent from '../views/Opponent.vue'
import { useGameStore } from '../stores/game'
import { GameState } from '../lib/messages'
import App from '../App.vue'
import { generateDeck } from './Hand.stories'

import './Opponent.stories.css'

const meta = {
  title: 'Opponent',
  component: Opponent,
  decorators: [() => ({
    template: '<div class="opponentArea"><story/></div>'
  })],
  tags: ['autodocs'],
  argTypes: {
    position: { control: 'number' },
  },
  args: { },
} satisfies Meta<typeof Opponent>

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
    isTurn: true,
    wonLastGame: false,
    connected: true,
    score: 0,
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
    }
  ]

  gameStore.winPlaces = []
}

export const Disconnected: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
    gameStore.opponents[0].connected = false
  }
}

export const WonFirstPlace: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
    gameStore.winPlaces = gameStore.opponents.concat(gameStore.self!)
  }
}

export const WonSecondPlace: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
    gameStore.winPlaces = [gameStore.self!].concat(gameStore.opponents)
  }
}

export const IsTurn: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
    gameStore.opponents[0].isTurn = true
  }
}

export const Waiting: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
  }
}

export const Passed: Story = {
  args: {
    position: 2,
  },
  play: async () => {
    reset()
    gameStore.opponents[0].isPassed = true
  }
}
