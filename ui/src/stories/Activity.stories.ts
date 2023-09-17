import type { Meta, StoryObj } from '@storybook/vue3'
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Activity from '../views/Activity.vue'
import { type EventRune, EventSeverity, Suit } from '../lib/models'
import { useGameStore } from '../stores/game'
import App from '../App.vue'

import './Activity.stories.css'

const meta = {
  title: 'Activity',
  component: Activity,
  decorators: [() => ({
    template: '<div class="activityArea"><story/></div>'
  })],
  args: { },
} satisfies Meta<typeof Activity>

export default meta
type Story = StoryObj<typeof meta>

const app = createApp(App)

app.use(createPinia())

const gameStore = useGameStore()

const pushEvent = ({
  severity,
  runes,
}: {
  severity: EventSeverity
  runes: EventRune[]
}) => {
  gameStore.events.push({
    severity,
    runes,
    timestamp: new Date(),
    toast: false,
  })
}

export const ErrorTextOnly: Story = {
  play: async () => {
    pushEvent({
      severity: EventSeverity.Error,
      runes: [
        {
          message: 'You joined the game',
        },
        {
          message: ', and you look great.',
        }
      ]
    })
  }
}

export const WarningTextOnly: Story = {
  play: async () => {
    pushEvent({
      severity: EventSeverity.Warning,
      runes: [
        {
          message: 'You joined the game',
        },
        {
          message: ', and you look great.',
        }
      ]
    })
  }
}

export const SuccessTextOnly: Story = {
  play: async () => {
    pushEvent({
      severity: EventSeverity.Success,
      runes: [
        {
          message: 'You joined the game',
        },
        {
          message: ', and you look great.',
        }
      ]
    })
  }
}

export const InfoTextAndRunes: Story = {
  play: async () => {
    pushEvent({
      severity: EventSeverity.Info,
      runes: [
        {
          message: 'You played ',
        },
        {
          card: {
            suit: Suit.Clubs,
            faceValue: 3,
            suitRank: 13,
            globalRank: 52,
          }
        },
        {
          card: {
            suit: Suit.Hearts,
            faceValue: 3,
            suitRank: 13,
            globalRank: 13,
          }
        }
      ]
    })
  }
}
