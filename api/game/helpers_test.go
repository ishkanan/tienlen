package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameCleaner(t *testing.T) {
	players := players{
		&player{Name: "Ben"},
	}
	outcomeMap := map[string]string{
		"":              "",
		"Aaron":         "Aaron",
		"Ben":           "",
		"Ben ":          "",
		" Daniel ":      "Daniel",
		"Antonietta":    "Antonietta",
		" Charlottina ": "Charlottin",
		"Tiến lên":      "",
	}
	for from, to := range outcomeMap {
		assert.Equal(t, to, cleanPlayerName(from, players, 10))
	}
}

func TestRandomNameGenerator(t *testing.T) {
	players := players{
		&player{Name: "Ben"},
	}
	name := theyWhoNotBeNamed(players, maxNameLength)
	assert.Equal(t, name, cleanPlayerName(name, players, maxNameLength))
}
