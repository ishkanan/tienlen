import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Controls from '../views/Controls.vue'
import { useGameStore } from '../stores/game'
import App from '../App.vue'

import './Controls.stories.css'

const meta = {
  title: 'Controls',
  component: Controls,
  decorators: [() => ({
    template: '<div class="controlsArea"><story/></div>'
  })],
  args: { },
} satisfies Meta<typeof Controls>

export default meta
type Story = StoryObj<typeof meta>

const app = createApp(App)

app.use(createPinia())

const gameStore = useGameStore()

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

export const DefaultView: Story = {}
