package hangman

import (
	"strings"

	"github.com/derailed/hangman/internal/game"
)

// Tally tracks the user score
type Tally struct {
	Status    game.Status `json:"status"`
	TurnsLeft int         `json:"turns_left"`
	Letters   string      `json:"letters"`
}

func tallyFromGame(g game.Game) Tally {
	return Tally{
		Status:    g.Status,
		TurnsLeft: g.TurnsLeft,
		Letters:   obfuscate(g.Letters, g.Guesses),
	}
}

func obfuscate(letters string, guesses []rune) string {
	res := make([]string, len(letters))
	for i, c := range letters {
		if alreadyGuessed(guesses, c) {
			res[i] = string(c)
		} else {
			res[i] = "_"
		}
	}
	return strings.Join(res, " ")
}

func alreadyGuessed(guesses []rune, l rune) bool {
	for _, c := range guesses {
		if c == l {
			return true
		}
	}
	return false
}
