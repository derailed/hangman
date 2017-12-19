package game

const (
	// Won indicates the game ended with a win
	Won = iota
	// Lost indicates the game ended with a loss
	Lost
	// Good indicates a correct guess
	Good
	// Bad indicates a bad guess
	Bad
	// Guessed indicates a letter was already guessed
	Guessed
	// Started indicates the game begun
	Started
)

type (
	// Status of game
	Status int

	// Guesser protocol to guess a letter for hangman
	Guesser interface {
		Guess(rune) *Game
	}

	// Game records state for hangman
	Game struct {
		Letters   string `json:"letters"`
		Status    Status `json:"status"`
		TurnsLeft int    `json:"turnsLeft"`
		Guesses   []rune `json:"guesses"`
	}
)

// NewGame starts a hangman game
func NewGame(word string) *Game {
	g := Game{Status: Started, TurnsLeft: 7, Letters: word, Guesses: []rune{}}
	return &g
}

// Guess a letter in the selected word
func (g *Game) Guess(l rune) *Game {
	if g.alreadyGuessed(l) {
		g.Status = Guessed
		return g
	}

	if g.inWord(l) {
		g.Guesses = append(g.Guesses, l)
		if g.isWon() {
			g.Status = Won
		} else {
			g.Status = Good
		}
		return g
	}

	g.Guesses = append(g.Guesses, l)
	g.TurnsLeft--
	if g.TurnsLeft == 0 {
		g.Status = Lost
	} else {
		g.Status = Bad
	}
	return g
}

func (g *Game) isWon() bool {
	for _, c := range g.Letters {
		if !g.alreadyGuessed(c) {
			return false
		}
	}
	return true
}

func (g *Game) alreadyGuessed(l rune) bool {
	for _, c := range g.Guesses {
		if c == l {
			return true
		}
	}
	return false
}

func (g *Game) inWord(l rune) bool {
	for _, c := range g.Letters {
		if c == l {
			return true
		}
	}
	return false
}
