package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// represents a player in the game
type player struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Position    int       `json:"position"` // 1,2,3,4
	Hand        []card    `json:"-"`
	CardsLeft   int       `json:"cardsLeft"`
	IsPassed    bool      `json:"isPassed"`
	IsTurn      bool      `json:"isTurn"`
	WonLastGame bool      `json:"wonLastGame"`
	Connected   bool      `json:"connected"`
}

// provides some helpers to help reduce clutter in game object
type players []*player

// GetByName gets a pointer to the player with matching name
func (p players) GetByName(name string) *player {
	for i := range p {
		if p[i].Name == name {
			return p[i]
		}
	}
	return nil
}

// DeleteByID remove a player by their ID
func (p players) DeleteByID(id uuid.UUID) players {
	kept := players{}
	for _, player := range p {
		if player.ID.String() != id.String() {
			kept = append(kept, player)
		}
	}
	return kept
}

// WithLowestCard returns the player with the lowest value card
func (p players) WithLowestCard() *player {
	lowestIndex := 0
	for i := range p {
		lowestCards := globalRankSort(p[lowestIndex].Hand)
		playerCards := globalRankSort(p[i].Hand)
		if playerCards[0].GlobalRank > lowestCards[0].GlobalRank {
			lowestIndex = i
		}
	}
	return p[lowestIndex]
}

// WonLastGame returns the player (if any) who won the last game
func (p players) WonLastGame() *player {
	for _, player := range p {
		if player.WonLastGame {
			return player
		}
	}
	return nil
}

// CurrentTurn returns the player who can play a hand or pass
func (p players) CurrentTurn() *player {
	for _, player := range p {
		if player.IsTurn {
			return player
		}
	}
	return nil // unreachable!
}

// AtPosition returns the player at the specified position
func (p players) AtPosition(position int) *player {
	for _, player := range p {
		if player.Position == position {
			return player
		}
	}
	return nil
}

// NextTurn returns the next player who can play a hand or pass
func (p players) NextTurn(after *player) *player {
	nextPosition := (after.Position % len(p)) + 1
	for {
		nextPlayer := p.AtPosition(nextPosition)
		if !nextPlayer.IsPassed {
			return nextPlayer
		}
		nextPosition = (nextPosition % len(p)) + 1
		if nextPosition == after.Position {
			// we've looped back around to the same player
			return after
		}
	}
}

// NextAvailablePosition returns the position for a newly-joined player
func (p players) NextAvailablePosition() int {
	for i := 1; i <= 4; i++ {
		if p.AtPosition(i) == nil {
			return i
		}
	}
	return 0 // unreachable!
}

var firstNames = []string{"Awesome", "Big", "Small", "Smart", "Good", "Great", "Adorable", "Fancy", "Witty", "Fast", "Eager", "Nice", "Lively", "Gifted", "Red", "Cute", "Clever", "Crazy", "Calm", "Cunning"}
var lastNames = []string{"Dog", "Cat", "Lion", "Eagle", "Bird", "Panda", "Fish", "Bear", "Hedgehog", "Quail", "Chicken", "Ant", "Bug", "Beetle", "Zebra", "Horse"}

// generates a new random name for a nameless player
func (p players) theyWhoNotBeNamed() string {
	rand.Seed(time.Now().UnixNano())
	for {
		name := fmt.Sprintf("%v %v",
			firstNames[rand.Intn(len(firstNames))],
			lastNames[rand.Intn(len(lastNames))],
		)
		if p.GetByName(name) == nil {
			return name
		}
	}
}
