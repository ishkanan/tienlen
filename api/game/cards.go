package game

import (
	"math/rand"
	"sort"
	"time"

	"github.com/meirf/gopart"
)

type pattern int
type suit int

const (
	patternSingle     pattern = 1 // 3
	patternDouble     pattern = 2 // 3,3
	patternTriple     pattern = 3 // and so on ...
	patternQuad       pattern = 4
	patternSeqSingles pattern = 5 // 3,4,5 ...
	patternSeqDoubles pattern = 6 // 3,3,4,4,5,5 ...
	patternSeqTriples pattern = 7 // and so on ...
	patternSeqQuads   pattern = 8
	patternInvalid    pattern = 9
	suitSpades        suit    = 1
	suitClubs         suit    = 2
	suitDiamonds      suit    = 3
	suitHearts        suit    = 4
)

// represents a playable card
type card struct {
	Suit       suit `json:"suit"`
	FaceValue  int  `json:"faceValue"`  // 1 = Ace, 13 = King
	SuitRank   int  `json:"suitRank"`   // 13 = lowest ("3"), 1 = highest ("2")
	GlobalRank int  `json:"globalRank"` // 52 = lowest ("3 Spades"), 1 = highest ("2 Hearts")
}

// returns a card stack sorted by suit rank, lowest to highest.
func suitRankSort(cards []card) []card {
	copied := append([]card(nil), cards...)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i].SuitRank > copied[j].SuitRank
	})
	return copied
}

// returns a card stack sorted by global rank, lowest to highest.
func globalRankSort(cards []card) []card {
	copied := append([]card(nil), cards...)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i].GlobalRank > copied[j].GlobalRank
	})
	return copied
}

// returns true if card set A ends with a higher card than set B
func areBetterCardsThan(setA, setB []card) bool {
	if beatenByChop(setA, setB) {
		return true
	}
	samePattern := determinePattern(setA) == determinePattern(setB)
	sameLength := len(setA) == len(setB)
	sortedA := globalRankSort(setA)
	sortedB := globalRankSort(setB)
	return samePattern && sameLength && sortedA[len(sortedA)-1].GlobalRank < sortedB[len(sortedB)-1].GlobalRank
}

// returns true if card set A is a chop and beats set B
func beatenByChop(setA, setB []card) bool {
	isSingleTwos := true
	for _, card := range setB {
		isSingleTwos = isSingleTwos && card.FaceValue == 2
	}
	patternA := determinePattern(setA)
	isChop := patternA == patternQuad || patternA == patternSeqDoubles
	if !isSingleTwos || !isChop {
		return false
	}
	chopLengthMap := map[int]int{1: 6, 2: 8, 3: 10, 4: 12}
	return (patternA == patternQuad && len(setB) == 1) || (len(setA) == chopLengthMap[len(setB)])
}

// returns the index the card with specified global rank appears in set of cards, -1 if not found
func cardInSet(globalRank int, cardSet []card) int {
	for i, card := range cardSet {
		if card.GlobalRank == globalRank {
			return i
		}
	}
	return -1
}

// returns the pattern (if any) of the cards.
func determinePattern(theCards []card) pattern {
	cards := suitRankSort(theCards)

	if len(cards) == 0 || len(cards) > 13 {
		return patternInvalid
	}

	// non-sequences (i.e. of a kind)
	if len(cards) <= 4 {
		identical := true
		previous := cards[0]
		for _, card := range cards {
			identical = identical && card.FaceValue == previous.FaceValue
			previous = card
		}
		if identical {
			return map[int]pattern{1: patternSingle, 2: patternDouble, 3: patternTriple, 4: patternQuad}[len(cards)]
		}
	}

	// sequences
	seqMap := map[int]pattern{1: patternSeqSingles, 2: patternSeqDoubles, 3: patternSeqTriples, 4: patternSeqQuads}
	for size, pattern := range seqMap {
		if len(cards)%size != 0 {
			continue
		}

		partitions := []gopart.IdxRange{}
		for partition := range gopart.Partition(len(cards), size) {
			partitions = append(partitions, partition)
		}
		if len(partitions) < 3 {
			continue
		}

		// three conditions:

		// 1) all cards in a partition must have same face value
		failed := false
		for _, partition := range partitions {
			previous := cards[partition.Low]
			for i := partition.Low; i < partition.High; i++ {
				failed = failed || cards[i].FaceValue != previous.FaceValue
				previous = cards[i]
			}
		}
		if failed {
			continue
		}

		// 2) suit_rank(partition N) - suit_rank(partition N-1) must equal -1
		previous := cards[partitions[0].Low]
		for i := 1; i < len(partitions); i++ {
			failed = failed || (cards[partitions[i].Low].SuitRank-previous.SuitRank) != -1
			previous = cards[partitions[i].Low]
		}
		if failed {
			continue
		}

		// 3) cannot end with 2's
		if globalRankSort(theCards)[len(theCards)-1].FaceValue == 2 {
			continue
		}

		return pattern
	}

	return patternInvalid
}

// returns a pseudo-shuffled deck of 52 cards
func buildShuffledDeck() []card {
	deck := make([]card, 0, 52)
	faces := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2}
	suits := []suit{suitSpades, suitClubs, suitDiamonds, suitHearts}
	globalRank := 52

	for i, face := range faces {
		for _, suit := range suits {
			deck = append(deck, card{
				Suit:       suit,
				FaceValue:  face,
				SuitRank:   13 - i,
				GlobalRank: globalRank,
			})
			globalRank--
		}
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled := make([]card, len(deck))
	for i, randIndex := range r.Perm(len(deck)) {
		shuffled[i] = deck[randIndex]
	}

	return shuffled
}
