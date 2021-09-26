package game

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// return a trimmed player name, or empty if player exists or invalid characters
func cleanPlayerName(name string, players players, maxLength int) string {
	name = strings.TrimSpace(name)
	if len(name) > maxLength {
		name = name[0:maxLength]
	}
	if players.GetByName(name) != nil {
		name = ""
	}
	r, _ := regexp.Compile("^[\u0020-\u007e\u00C0-\u00ff]+$")
	matches := r.FindAllString(name, -1)
	if matches == nil || len(matches) != 1 {
		name = ""
	}
	return name
}

// generates a new, unique random player name
func theyWhoNotBeNamed(players players, maxLength int) string {
	rand.Seed(time.Now().UnixNano())
	for {
		name := fmt.Sprintf("%v %v",
			firstNames[rand.Intn(len(firstNames))],
			lastNames[rand.Intn(len(lastNames))],
		)
		if cleanPlayerName(name, players, maxLength) != "" {
			return name
		}
	}
}
