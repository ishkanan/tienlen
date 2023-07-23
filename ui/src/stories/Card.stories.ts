import type { Meta, StoryObj } from '@storybook/vue3'

import Card from '../components/Card.vue'
import { Suit } from '../lib/models'

const meta = {
  title: 'Card',
  component: Card,
  tags: ['autodocs'],
  argTypes: {
    card: { control: 'object' },
    selectable: { control: 'boolean' },
    showFace: { control: 'boolean' },
    onSelected: { action: 'clicked' },
  },
  args: { },
} satisfies Meta<typeof Card>

export default meta
type Story = StoryObj<typeof meta>

export const HighestPlayable: Story = {
  args: {
    card: {
      suit: Suit.Hearts,
      faceValue: 2,
      suitRank: 1,
      globalRank: 1,
    },
    selectable: true,
    showFace: true,
  },
};

export const LowestPlayable: Story = {
  args: {
    card: {
      suit: Suit.Clubs,
      faceValue: 3,
      suitRank: 13,
      globalRank: 52,
    },
    selectable: true,
    showFace: true,
  },
};

export const UnplayableBlack: Story = {
  args: {
    card: {
      suit: Suit.Clubs,
      faceValue: 0,
      suitRank: 0,
      globalRank: 0,
    },
    selectable: false,
    showFace: false,
  },
};

export const UnplayableRed: Story = {
  args: {
    card: {
      suit: Suit.Diamonds,
      faceValue: 0,
      suitRank: 0,
      globalRank: 0,
    },
    selectable: false,
    showFace: false,
  },
};
