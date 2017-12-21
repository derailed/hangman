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
		Letters:   obfuscate(g),
	}
}

func obfuscate(g game.Game) string {
	res := make([]string, len(g.Letters))
	for i, c := range g.Letters {
		if g.AlreadyGuessed(c) {
			res[i] = string(c)
		} else {
			res[i] = "_"
		}
	}
	return strings.Join(res, " ")
}
