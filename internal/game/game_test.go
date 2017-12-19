package game_test

import (
	"testing"

	"github.com/derailed/hangman2/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestGoodGuess(t *testing.T) {
	g := game.NewGame("blee")
	g = g.Guess('e')

	assert.Equal(t, game.Status(game.Good), g.Status)
	assert.Equal(t, 7, g.TurnsLeft)
	assert.Equal(t, "blee", g.Letters)
	assert.Equal(t, []rune{'e'}, g.Guesses)
}

func TestBadGuess(t *testing.T) {
	g := game.NewGame("blee")
	g = g.Guess('z')

	assert.Equal(t, game.Status(game.Bad), g.Status)
	assert.Equal(t, 6, g.TurnsLeft)
	assert.Equal(t, "blee", g.Letters)
	assert.Equal(t, []rune{'z'}, g.Guesses)
}

func TestAlreadyGuessed(t *testing.T) {
	g := game.NewGame("blee")
	g = g.Guess('b')
	g = g.Guess('b')

	assert.Equal(t, game.Status(game.Guessed), g.Status)
	assert.Equal(t, 7, g.TurnsLeft)
	assert.Equal(t, "blee", g.Letters)
	assert.Equal(t, []rune{'b'}, g.Guesses)
}

func TestWin(t *testing.T) {
	guesses := []struct {
		g rune
		s game.Status
	}{
		{g: 'b', s: game.Good},
		{g: 'l', s: game.Good},
		{g: 'e', s: game.Won},
	}

	g := game.NewGame("blee")
	for _, r := range guesses {
		g = g.Guess(r.g)
		assert.Equal(t, r.s, g.Status)
	}
}
func TestLoose(t *testing.T) {
	guesses := []struct {
		g rune
		s game.Status
	}{
		{g: 'q', s: game.Bad},
		{g: 'r', s: game.Bad},
		{g: 's', s: game.Bad},
		{g: 'u', s: game.Bad},
		{g: 'v', s: game.Bad},
		{g: 'x', s: game.Bad},
		{g: 'z', s: game.Lost},
	}

	g := game.NewGame("blee")
	for _, r := range guesses {
		g = g.Guess(r.g)
		assert.Equal(t, r.s, g.Status)
	}
}
