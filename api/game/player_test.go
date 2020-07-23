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

func TestNextTurnClockwiseNoWins(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{card{}}},
		&player{Position: 2, Hand: []card{card{}}},
		&player{Position: 3, Hand: []card{card{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{card{}}},
	}
	assert.Equal(t, 2, p.NextTurn(p[0], true).Position)
	assert.Equal(t, 4, p.NextTurn(p[1], true).Position)
	assert.Equal(t, 4, p.NextTurn(p[2], true).Position)
	assert.Equal(t, 1, p.NextTurn(p[3], true).Position)
}

func TestNextTurnClockwiseOneWinner(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{card{}}},
		&player{Position: 2, Hand: []card{}},
		&player{Position: 3, Hand: []card{card{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{card{}}},
	}
	assert.Equal(t, 4, p.NextTurn(p[0], true).Position)
	assert.Equal(t, 4, p.NextTurn(p[1], true).Position)
	assert.Equal(t, 4, p.NextTurn(p[2], true).Position)
	assert.Equal(t, 1, p.NextTurn(p[3], true).Position)
}

func TestNextTurnAntiClockwiseNoWins(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{card{}}},
		&player{Position: 2, Hand: []card{card{}}},
		&player{Position: 3, Hand: []card{card{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{card{}}},
	}
	assert.Equal(t, 2, p.NextTurn(p[3], false).Position)
	assert.Equal(t, 2, p.NextTurn(p[2], false).Position)
	assert.Equal(t, 1, p.NextTurn(p[1], false).Position)
	assert.Equal(t, 4, p.NextTurn(p[0], false).Position)
}

func TestNextTurnAntiClockwiseOneWinner(t *testing.T) {
	p := players{
		&player{Position: 1, Hand: []card{card{}}},
		&player{Position: 2, Hand: []card{}},
		&player{Position: 3, Hand: []card{card{}}, IsPassed: true},
		&player{Position: 4, Hand: []card{card{}}},
	}
	assert.Equal(t, 1, p.NextTurn(p[3], false).Position)
	assert.Equal(t, 1, p.NextTurn(p[2], false).Position)
	assert.Equal(t, 1, p.NextTurn(p[1], false).Position)
	assert.Equal(t, 4, p.NextTurn(p[0], false).Position)
}
