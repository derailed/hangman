package game

import (
	"strings"
)

const (
	Won = iota
	Lost
	Good
	Bad
	Guessed
	Started
)

type (
	Status int

	Guesser interface {
		Guess(rune) (*Game, Tally)
	}

	Game struct {
		Letters   string `json:"letter"`
		Status    Status `json:"status"`
		TurnsLeft int    `json:"turns_left`
		Guesses   []rune `json:"guesses`
	}

	Tally struct {
		Status    Status `json:"status"`
		TurnsLeft int    `json:"turns_left`
		Letters   string `json:"letters`
	}
)

// NewGame of hangman
func NewGame(word string) *Game {
	// Need to talk to word list and get rand word
	g := Game{Status: Started, TurnsLeft: 7, Letters: word}

	return &g
}

func (g *Game) returnWithTally() (*Game, Tally) {
	return g, g.tallyFromGame()
}

// Guess a letter in the selected word
func (g *Game) Guess(l rune) (*Game, Tally) {
	if g.alreadyGuessed(l) {
		g.Status = Guessed
		return g.returnWithTally()
	}

	if g.inWord(l) {
		g.Guesses = append(g.Guesses, l)
		if g.isWon() {
			g.Status = Won
		} else {
			g.Status = Good
		}
		return g.returnWithTally()
	}

	g.Guesses = append(g.Guesses, l)
	g.TurnsLeft--
	if g.TurnsLeft == 0 {
		g.Status = Lost
	} else {
		g.Status = Bad
	}
	return g.returnWithTally()
}

func (g *Game) isWon() bool {
	for _, c := range g.Letters {
		if !g.alreadyGuessed(c) {
			return false
		}
	}
	return true
}

func (g *Game) tallyFromGame() Tally {
	return Tally{
		Status:    g.Status,
		TurnsLeft: g.TurnsLeft,
		Letters:   g.obfuscate(),
	}
}

func (g *Game) obfuscate() string {
	res := make([]string, len(g.Letters))
	for i, c := range g.Letters {
		if g.alreadyGuessed(c) {
			res[i] = string(c)
		} else {
			res[i] = "_"
		}
	}
	return strings.Join(res, " ")
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
