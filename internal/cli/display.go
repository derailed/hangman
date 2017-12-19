package cli

import (
	"fmt"

	"github.com/derailed/hangman2/internal/game"
)

func Display(tally game.Tally) {
	fmt.Printf("\n%10s: %s\n", "Letters", tally.Letters)
	fmt.Printf("%10s: %s [%d]\n", "Status", statusToH(tally.Status), tally.TurnsLeft)
}

func statusToH(s game.Status) string {
	switch s {
	case game.Guessed:
		return "Already Guessed"
	case game.Won:
		return "Won"
	case game.Lost:
		return "Lost"
	case game.Good:
		return "Good"
	case game.Bad:
		return "Bad"
	default:
		return fmt.Sprintf("Initialized")
	}
}
