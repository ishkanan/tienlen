package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextAvailablePosition(t *testing.T) {
	p := players{}
	assert.Equal(t, 1, p.NextAvailablePosition())
	p = append(p, &player{Position: 1})
	assert.Equal(t, 2, p.NextAvailablePosition())
	p = append(p, &player{Position: 2})
	assert.Equal(t, 3, p.NextAvailablePosition())
	p = append(p, &player{Position: 3})
	assert.Equal(t, 4, p.NextAvailablePosition())
	p = append(p, &player{Position: 4})
	assert.Equal(t, 0, p.NextAvailablePosition())
}

func TestNextTurnNoWins(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{{}}},
		&player{Position: 2, Hand: []card{{}}},
		&player{Position: 3, Hand: []card{{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{{}}},
	}
	assert.Equal(t, 2, p.NextTurn(p[0]).Position)
	assert.Equal(t, 4, p.NextTurn(p[1]).Position)
	assert.Equal(t, 4, p.NextTurn(p[2]).Position)
	assert.Equal(t, 1, p.NextTurn(p[3]).Position)
}

func TestNextTurnOneWinner(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{{}}},
		&player{Position: 2, Hand: []card{}},
		&player{Position: 3, Hand: []card{{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{{}}},
	}
	assert.Equal(t, 4, p.NextTurn(p[0]).Position)
	assert.Equal(t, 4, p.NextTurn(p[1]).Position)
	assert.Equal(t, 4, p.NextTurn(p[2]).Position)
	assert.Equal(t, 1, p.NextTurn(p[3]).Position)
}

func TestNextTurnTwoWinners(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{}},
		&player{Position: 2, Hand: []card{}},
		&player{Position: 3, Hand: []card{{}}},
		&player{Position: 4, Hand: []card{{}}},
	}
	assert.Equal(t, 3, p.NextTurn(p[0]).Position)
	assert.Equal(t, 3, p.NextTurn(p[1]).Position)
	assert.Equal(t, 4, p.NextTurn(p[2]).Position)
	assert.Equal(t, 3, p.NextTurn(p[3]).Position)
}

func TestGetByName(t *testing.T) {
	p := players{
		&player{Name: "p1"},
		&player{Name: "p2"},
	}
	assert.Equal(t, p[0], p.GetByName("p1"))
	assert.Equal(t, p[1], p.GetByName("p2"))
	assert.Equal(t, (*player)(nil), p.GetByName("p3"))
}

func TestDeleteDisconnected(t *testing.T) {
	p := players{
		&player{Name: "p1", Connected: true},
		&player{Name: "p2"},
	}
	assert.Equal(t, 1, len(p.DeleteDisconnected()))
	assert.Equal(t, players{p[0]}, p.DeleteDisconnected())
}

func TestDeleteByName(t *testing.T) {
	p := players{
		&player{Name: "p1"},
		&player{Name: "p2"},
	}
	assert.Equal(t, 1, len(p.DeleteByName("p1")))
	assert.Equal(t, players{p[1]}, p.DeleteByName("p1"))
}

func TestPassedAndPlacedCount(t *testing.T) {
	p := players{
		&player{Name: "p1", IsPassed: true},
		&player{Name: "p2", CardsLeft: 1},
		&player{Name: "p3"},
	}
	assert.Equal(t, 2, p.PassedAndPlacedCount())
}

func TestWithLowestCard(t *testing.T) {
	p := players{
		&player{Name: "p1", Hand: []card{
			{GlobalRank: 36}, {GlobalRank: 20},
		}},
		&player{Name: "p2", Hand: []card{
			{GlobalRank: 1}, {GlobalRank: 26},
		}},
		&player{Name: "p3", Hand: []card{
			{GlobalRank: 52}, {GlobalRank: 42},
		}},
	}
	assert.Equal(t, p[2], p.WithLowestCard())

	p = players{
		&player{Name: "p1", Hand: []card{
			{GlobalRank: 36}, {GlobalRank: 51},
		}},
		&player{Name: "p2", Hand: []card{
			{GlobalRank: 1}, {GlobalRank: 26},
		}},
		&player{Name: "p3", Hand: []card{
			{GlobalRank: 2}, {GlobalRank: 42},
		}},
	}
	assert.Equal(t, p[0], p.WithLowestCard())
}

func TestWonLastGame(t *testing.T) {
	p := players{
		&player{Name: "p1", WonLastGame: true},
		&player{Name: "p2"},
		&player{Name: "p3"},
	}
	assert.Equal(t, p[0], p.WonLastGame())
}

func TestCurrentTurn(t *testing.T) {
	p := players{
		&player{Name: "p1", IsTurn: true},
		&player{Name: "p2"},
		&player{Name: "p3"},
	}
	assert.Equal(t, p[0], p.CurrentTurn())
}

func TestAtPosition(t *testing.T) {
	p := players{
		&player{Name: "p1", Position: 1},
		&player{Name: "p2", Position: 2},
		&player{Name: "p3", Position: 3},
	}
	assert.Equal(t, p[0], p.AtPosition(1))
	assert.Equal(t, p[1], p.AtPosition(2))
	assert.Equal(t, p[2], p.AtPosition(3))
	assert.Equal(t, (*player)(nil), p.AtPosition(4))
}
