package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffleDeck(t *testing.T) {
	deck := buildShuffledDeck()

	assert.Equal(t, 52, len(deck))

	// check face values per suit
	expectedDeck := []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		card{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		card{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		card{Suit: suitSpades, FaceValue: 7, SuitRank: 9, GlobalRank: 36},
		card{Suit: suitSpades, FaceValue: 8, SuitRank: 8, GlobalRank: 32},
		card{Suit: suitSpades, FaceValue: 9, SuitRank: 7, GlobalRank: 28},
		card{Suit: suitSpades, FaceValue: 10, SuitRank: 6, GlobalRank: 24},
		card{Suit: suitSpades, FaceValue: 11, SuitRank: 5, GlobalRank: 20},
		card{Suit: suitSpades, FaceValue: 12, SuitRank: 4, GlobalRank: 16},
		card{Suit: suitSpades, FaceValue: 13, SuitRank: 3, GlobalRank: 12},
		card{Suit: suitSpades, FaceValue: 1, SuitRank: 2, GlobalRank: 8},
		card{Suit: suitSpades, FaceValue: 2, SuitRank: 1, GlobalRank: 4},

		card{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		card{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		card{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		card{Suit: suitClubs, FaceValue: 6, SuitRank: 10, GlobalRank: 39},
		card{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		card{Suit: suitClubs, FaceValue: 8, SuitRank: 8, GlobalRank: 31},
		card{Suit: suitClubs, FaceValue: 9, SuitRank: 7, GlobalRank: 27},
		card{Suit: suitClubs, FaceValue: 10, SuitRank: 6, GlobalRank: 23},
		card{Suit: suitClubs, FaceValue: 11, SuitRank: 5, GlobalRank: 19},
		card{Suit: suitClubs, FaceValue: 12, SuitRank: 4, GlobalRank: 15},
		card{Suit: suitClubs, FaceValue: 13, SuitRank: 3, GlobalRank: 11},
		card{Suit: suitClubs, FaceValue: 1, SuitRank: 2, GlobalRank: 7},
		card{Suit: suitClubs, FaceValue: 2, SuitRank: 1, GlobalRank: 3},

		card{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		card{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		card{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		card{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
		card{Suit: suitDiamonds, FaceValue: 7, SuitRank: 9, GlobalRank: 34},
		card{Suit: suitDiamonds, FaceValue: 8, SuitRank: 8, GlobalRank: 30},
		card{Suit: suitDiamonds, FaceValue: 9, SuitRank: 7, GlobalRank: 26},
		card{Suit: suitDiamonds, FaceValue: 10, SuitRank: 6, GlobalRank: 22},
		card{Suit: suitDiamonds, FaceValue: 11, SuitRank: 5, GlobalRank: 18},
		card{Suit: suitDiamonds, FaceValue: 12, SuitRank: 4, GlobalRank: 14},
		card{Suit: suitDiamonds, FaceValue: 13, SuitRank: 3, GlobalRank: 10},
		card{Suit: suitDiamonds, FaceValue: 1, SuitRank: 2, GlobalRank: 6},
		card{Suit: suitDiamonds, FaceValue: 2, SuitRank: 1, GlobalRank: 2},

		card{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
		card{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		card{Suit: suitHearts, FaceValue: 5, SuitRank: 11, GlobalRank: 41},
		card{Suit: suitHearts, FaceValue: 6, SuitRank: 10, GlobalRank: 37},
		card{Suit: suitHearts, FaceValue: 7, SuitRank: 9, GlobalRank: 33},
		card{Suit: suitHearts, FaceValue: 8, SuitRank: 8, GlobalRank: 29},
		card{Suit: suitHearts, FaceValue: 9, SuitRank: 7, GlobalRank: 25},
		card{Suit: suitHearts, FaceValue: 10, SuitRank: 6, GlobalRank: 21},
		card{Suit: suitHearts, FaceValue: 11, SuitRank: 5, GlobalRank: 17},
		card{Suit: suitHearts, FaceValue: 12, SuitRank: 4, GlobalRank: 13},
		card{Suit: suitHearts, FaceValue: 13, SuitRank: 3, GlobalRank: 9},
		card{Suit: suitHearts, FaceValue: 1, SuitRank: 2, GlobalRank: 5},
		card{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	expectedDeck = globalRankSort(expectedDeck)
	deck = globalRankSort(deck)
	for i := 0; i < 52; i++ {
		assert.Equal(t, expectedDeck[i], deck[i], "%+v", deck[i])
	}
}

func TestPatternCheck(t *testing.T) {
	cards := []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		card{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		card{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
	}
	assert.Equal(t, patternSingle, determinePattern([]card{cards[0]}))
	assert.Equal(t, patternDouble, determinePattern(cards[0:2]))
	assert.Equal(t, patternTriple, determinePattern(cards[0:3]))
	assert.Equal(t, patternQuad, determinePattern(cards[0:4]))

	cards = []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		card{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		card{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		card{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		card{Suit: suitClubs, FaceValue: 8, SuitRank: 8, GlobalRank: 31},
		card{Suit: suitClubs, FaceValue: 9, SuitRank: 7, GlobalRank: 27},
		card{Suit: suitClubs, FaceValue: 10, SuitRank: 6, GlobalRank: 23},
		card{Suit: suitDiamonds, FaceValue: 11, SuitRank: 5, GlobalRank: 18},
		card{Suit: suitDiamonds, FaceValue: 12, SuitRank: 4, GlobalRank: 14},
		card{Suit: suitDiamonds, FaceValue: 13, SuitRank: 3, GlobalRank: 10},
		card{Suit: suitHearts, FaceValue: 1, SuitRank: 2, GlobalRank: 5},
	}
	for i := 3; i < 12; i++ {
		assert.Equal(t, patternSeqSingles, determinePattern(cards[0:i]))
	}

	cards = []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		card{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		card{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		card{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		card{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		card{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
		card{Suit: suitHearts, FaceValue: 6, SuitRank: 10, GlobalRank: 37},
		card{Suit: suitSpades, FaceValue: 7, SuitRank: 9, GlobalRank: 36},
		card{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		card{Suit: suitDiamonds, FaceValue: 8, SuitRank: 8, GlobalRank: 30},
		card{Suit: suitHearts, FaceValue: 8, SuitRank: 8, GlobalRank: 29},
	}
	for i := 6; i < 12; i += 2 {
		assert.Equal(t, patternSeqDoubles, determinePattern(cards[0:i]))
	}

	cards = []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		card{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		card{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		card{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		card{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		card{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		card{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		card{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		card{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		card{Suit: suitClubs, FaceValue: 6, SuitRank: 10, GlobalRank: 39},
		card{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
	}
	assert.Equal(t, patternSeqTriples, determinePattern(cards[0:9]))
	assert.Equal(t, patternSeqTriples, determinePattern(cards[0:12]))

	cards = []card{
		card{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		card{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		card{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		card{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
		card{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		card{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		card{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		card{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		card{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		card{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		card{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		card{Suit: suitHearts, FaceValue: 5, SuitRank: 11, GlobalRank: 41},
	}
	assert.Equal(t, patternSeqQuads, determinePattern(cards))
}

func TestGlobalRankSort(t *testing.T) {
	cards := []card{
		card{GlobalRank: 1},
		card{GlobalRank: 33},
		card{GlobalRank: 45},
		card{GlobalRank: 27},
		card{GlobalRank: 52},
	}
	sorted := globalRankSort(cards)
	assert.Equal(t, 52, sorted[0].GlobalRank)
	assert.Equal(t, 45, sorted[1].GlobalRank)
	assert.Equal(t, 33, sorted[2].GlobalRank)
	assert.Equal(t, 27, sorted[3].GlobalRank)
	assert.Equal(t, 1, sorted[4].GlobalRank)
}

func TestSuitRankSort(t *testing.T) {
	cards := []card{
		card{SuitRank: 1},
		card{SuitRank: 13},
		card{SuitRank: 13},
		card{SuitRank: 7},
		card{SuitRank: 4},
	}
	sorted := suitRankSort(cards)
	assert.Equal(t, 13, sorted[0].SuitRank)
	assert.Equal(t, 13, sorted[1].SuitRank)
	assert.Equal(t, 7, sorted[2].SuitRank)
	assert.Equal(t, 4, sorted[3].SuitRank)
	assert.Equal(t, 1, sorted[4].SuitRank)
}
