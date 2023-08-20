import type { Meta, StoryObj } from '@storybook/vue3'

import Hand from '../components/Hand.vue'
import { Suit, type Card } from '../lib/models'

import './Card.stories.css'

const meta = {
  title: 'Hand',
  component: Hand,
  tags: ['autodocs'],
  argTypes: {
    cards: { control: 'object' },
    onSelected: { action: 'clicked' },
  },
  args: { },
} satisfies Meta<typeof Hand>

export default meta
type Story = StoryObj<typeof meta>

const generateDeck = (): Card[] => {
	const faces = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2]
	const suits = [Suit.Spades, Suit.Clubs, Suit.Diamonds, Suit.Hearts]
	let deck: Card[] = []
	let globalRank = 52

	faces.forEach((face, faceIndex) => {
		suits.forEach(suit => {
			deck.push({
				suit,
				faceValue: face,
				suitRank: 13 - faceIndex,
				globalRank,
			})
			globalRank--
		})
	})

  const randInt = (min: number, max: number) => (
    Math.floor(Math.random() * (max - min + 1) ) + min
  )

	let shuffled: Card[] = []
	for(let i = 0; i < 13; i++) {
    let cardIndex = randInt(0, deck.length - 1)
		shuffled.push(deck[cardIndex])
    deck.splice(cardIndex, 1)
	}

	return shuffled
}

export const FullHand: Story = {
  args: {
    cards: generateDeck().slice(0, 13),
  },
}

export const SmallHand: Story = {
  args: {
    cards: generateDeck().slice(0, 5),
  },
}
