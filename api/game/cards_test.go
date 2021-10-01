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
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		{Suit: suitSpades, FaceValue: 7, SuitRank: 9, GlobalRank: 36},
		{Suit: suitSpades, FaceValue: 8, SuitRank: 8, GlobalRank: 32},
		{Suit: suitSpades, FaceValue: 9, SuitRank: 7, GlobalRank: 28},
		{Suit: suitSpades, FaceValue: 10, SuitRank: 6, GlobalRank: 24},
		{Suit: suitSpades, FaceValue: 11, SuitRank: 5, GlobalRank: 20},
		{Suit: suitSpades, FaceValue: 12, SuitRank: 4, GlobalRank: 16},
		{Suit: suitSpades, FaceValue: 13, SuitRank: 3, GlobalRank: 12},
		{Suit: suitSpades, FaceValue: 1, SuitRank: 2, GlobalRank: 8},
		{Suit: suitSpades, FaceValue: 2, SuitRank: 1, GlobalRank: 4},

		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		{Suit: suitClubs, FaceValue: 6, SuitRank: 10, GlobalRank: 39},
		{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		{Suit: suitClubs, FaceValue: 8, SuitRank: 8, GlobalRank: 31},
		{Suit: suitClubs, FaceValue: 9, SuitRank: 7, GlobalRank: 27},
		{Suit: suitClubs, FaceValue: 10, SuitRank: 6, GlobalRank: 23},
		{Suit: suitClubs, FaceValue: 11, SuitRank: 5, GlobalRank: 19},
		{Suit: suitClubs, FaceValue: 12, SuitRank: 4, GlobalRank: 15},
		{Suit: suitClubs, FaceValue: 13, SuitRank: 3, GlobalRank: 11},
		{Suit: suitClubs, FaceValue: 1, SuitRank: 2, GlobalRank: 7},
		{Suit: suitClubs, FaceValue: 2, SuitRank: 1, GlobalRank: 3},

		{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
		{Suit: suitDiamonds, FaceValue: 7, SuitRank: 9, GlobalRank: 34},
		{Suit: suitDiamonds, FaceValue: 8, SuitRank: 8, GlobalRank: 30},
		{Suit: suitDiamonds, FaceValue: 9, SuitRank: 7, GlobalRank: 26},
		{Suit: suitDiamonds, FaceValue: 10, SuitRank: 6, GlobalRank: 22},
		{Suit: suitDiamonds, FaceValue: 11, SuitRank: 5, GlobalRank: 18},
		{Suit: suitDiamonds, FaceValue: 12, SuitRank: 4, GlobalRank: 14},
		{Suit: suitDiamonds, FaceValue: 13, SuitRank: 3, GlobalRank: 10},
		{Suit: suitDiamonds, FaceValue: 1, SuitRank: 2, GlobalRank: 6},
		{Suit: suitDiamonds, FaceValue: 2, SuitRank: 1, GlobalRank: 2},

		{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
		{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		{Suit: suitHearts, FaceValue: 5, SuitRank: 11, GlobalRank: 41},
		{Suit: suitHearts, FaceValue: 6, SuitRank: 10, GlobalRank: 37},
		{Suit: suitHearts, FaceValue: 7, SuitRank: 9, GlobalRank: 33},
		{Suit: suitHearts, FaceValue: 8, SuitRank: 8, GlobalRank: 29},
		{Suit: suitHearts, FaceValue: 9, SuitRank: 7, GlobalRank: 25},
		{Suit: suitHearts, FaceValue: 10, SuitRank: 6, GlobalRank: 21},
		{Suit: suitHearts, FaceValue: 11, SuitRank: 5, GlobalRank: 17},
		{Suit: suitHearts, FaceValue: 12, SuitRank: 4, GlobalRank: 13},
		{Suit: suitHearts, FaceValue: 13, SuitRank: 3, GlobalRank: 9},
		{Suit: suitHearts, FaceValue: 1, SuitRank: 2, GlobalRank: 5},
		{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	expectedDeck = globalRankSort(expectedDeck)
	deck = globalRankSort(deck)
	for i := 0; i < 52; i++ {
		assert.Equal(t, expectedDeck[i], deck[i], "%+v", deck[i])
	}
}

func TestPatternCheck(t *testing.T) {
	cards := []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
	}
	assert.Equal(t, patternSingle, determinePattern([]card{cards[0]}))
	assert.Equal(t, patternDouble, determinePattern(cards[0:2]))
	assert.Equal(t, patternTriple, determinePattern(cards[0:3]))
	assert.Equal(t, patternQuad, determinePattern(cards[0:4]))

	cards = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		{Suit: suitClubs, FaceValue: 8, SuitRank: 8, GlobalRank: 31},
		{Suit: suitClubs, FaceValue: 9, SuitRank: 7, GlobalRank: 27},
		{Suit: suitClubs, FaceValue: 10, SuitRank: 6, GlobalRank: 23},
		{Suit: suitDiamonds, FaceValue: 11, SuitRank: 5, GlobalRank: 18},
		{Suit: suitDiamonds, FaceValue: 12, SuitRank: 4, GlobalRank: 14},
		{Suit: suitDiamonds, FaceValue: 13, SuitRank: 3, GlobalRank: 10},
		{Suit: suitHearts, FaceValue: 1, SuitRank: 2, GlobalRank: 5},
	}
	for i := 3; i < 12; i++ {
		assert.Equal(t, patternSeqSingles, determinePattern(cards[0:i]))
	}

	cards = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
		{Suit: suitHearts, FaceValue: 6, SuitRank: 10, GlobalRank: 37},
		{Suit: suitSpades, FaceValue: 7, SuitRank: 9, GlobalRank: 36},
		{Suit: suitClubs, FaceValue: 7, SuitRank: 9, GlobalRank: 35},
		{Suit: suitDiamonds, FaceValue: 8, SuitRank: 8, GlobalRank: 30},
		{Suit: suitHearts, FaceValue: 8, SuitRank: 8, GlobalRank: 29},
	}
	for i := 6; i < 12; i += 2 {
		assert.Equal(t, patternSeqDoubles, determinePattern(cards[0:i]))
	}

	cards = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
		{Suit: suitClubs, FaceValue: 6, SuitRank: 10, GlobalRank: 39},
		{Suit: suitDiamonds, FaceValue: 6, SuitRank: 10, GlobalRank: 38},
	}
	assert.Equal(t, patternSeqTriples, determinePattern(cards[0:9]))
	assert.Equal(t, patternSeqTriples, determinePattern(cards[0:12]))

	cards = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 50},
		{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 49},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitClubs, FaceValue: 4, SuitRank: 12, GlobalRank: 47},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		{Suit: suitDiamonds, FaceValue: 5, SuitRank: 11, GlobalRank: 42},
		{Suit: suitHearts, FaceValue: 5, SuitRank: 11, GlobalRank: 41},
	}
	assert.Equal(t, patternSeqQuads, determinePattern(cards))
}

func TestGlobalRankSort(t *testing.T) {
	cards := []card{
		{GlobalRank: 1},
		{GlobalRank: 33},
		{GlobalRank: 45},
		{GlobalRank: 27},
		{GlobalRank: 52},
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
		{SuitRank: 1},
		{SuitRank: 13},
		{SuitRank: 13},
		{SuitRank: 7},
		{SuitRank: 4},
	}
	sorted := suitRankSort(cards)
	assert.Equal(t, 13, sorted[0].SuitRank)
	assert.Equal(t, 13, sorted[1].SuitRank)
	assert.Equal(t, 7, sorted[2].SuitRank)
	assert.Equal(t, 4, sorted[3].SuitRank)
	assert.Equal(t, 1, sorted[4].SuitRank)
}

func TestBeatenByChop(t *testing.T) {
	// nonsense
	hand1 := []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
	}
	hand2 := []card{
		{Suit: suitDiamonds, FaceValue: 8, SuitRank: 8, GlobalRank: 30},
		{Suit: suitHearts, FaceValue: 8, SuitRank: 8, GlobalRank: 29},
	}
	assert.Equal(t, false, beatenByChop(hand1, hand2))
	assert.Equal(t, false, beatenByChop(hand2, hand1))

	// quad and single 2
	hand1 = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 39},
		{Suit: suitDiamonds, FaceValue: 3, SuitRank: 13, GlobalRank: 26},
		{Suit: suitHearts, FaceValue: 3, SuitRank: 13, GlobalRank: 13},
	}
	hand2 = []card{
		{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	assert.Equal(t, true, beatenByChop(hand1, hand2))
	assert.Equal(t, false, beatenByChop(hand2, hand1))

	// quad and double 2
	hand1 = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitSpades, FaceValue: 4, SuitRank: 12, GlobalRank: 48},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 40},
	}
	hand2 = []card{
		{Suit: suitDiamonds, FaceValue: 2, SuitRank: 1, GlobalRank: 14},
		{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	assert.Equal(t, false, beatenByChop(hand1, hand2))
	assert.Equal(t, false, beatenByChop(hand2, hand1))

	// doubles seq and single 2
	hand1 = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
	}
	hand2 = []card{
		{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	assert.Equal(t, true, beatenByChop(hand1, hand2))
	assert.Equal(t, false, beatenByChop(hand2, hand1))

	// doubles seq and double 2
	hand1 = []card{
		{Suit: suitSpades, FaceValue: 3, SuitRank: 13, GlobalRank: 52},
		{Suit: suitClubs, FaceValue: 3, SuitRank: 13, GlobalRank: 51},
		{Suit: suitDiamonds, FaceValue: 4, SuitRank: 12, GlobalRank: 46},
		{Suit: suitHearts, FaceValue: 4, SuitRank: 12, GlobalRank: 45},
		{Suit: suitSpades, FaceValue: 5, SuitRank: 11, GlobalRank: 44},
		{Suit: suitClubs, FaceValue: 5, SuitRank: 11, GlobalRank: 43},
		{Suit: suitSpades, FaceValue: 6, SuitRank: 10, GlobalRank: 42},
		{Suit: suitClubs, FaceValue: 6, SuitRank: 10, GlobalRank: 41},
	}
	hand2 = []card{
		{Suit: suitDiamonds, FaceValue: 2, SuitRank: 1, GlobalRank: 14},
		{Suit: suitHearts, FaceValue: 2, SuitRank: 1, GlobalRank: 1},
	}
	assert.Equal(t, true, beatenByChop(hand1, hand2))
	assert.Equal(t, false, beatenByChop(hand2, hand1))
}

func TestCardInSet(t *testing.T) {
	cards := []card{
		{GlobalRank: 52},
		{GlobalRank: 24},
	}
	assert.Equal(t, 0, cardInSet(52, cards))
	assert.Equal(t, 1, cardInSet(24, cards))
	assert.Equal(t, -1, cardInSet(51, cards))
}
