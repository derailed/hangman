package game_test

import (
	"testing"

	"github.com/derailed/hangman2/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestGoodGuess(t *testing.T) {
	game := game.NewGame("blee")
	game, tally := game.Guess('e')

	assert.Equal(t, game.Status(game.Good), tally.Status)
	assert.Equal(t, 7, tally.TurnsLeft)
	assert.Equal(t, "_ _ e e", tally.Letters)

	assert.Equal(t, game.Status(game.Good), game.Status)
	assert.Equal(t, 7, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'e'}, game.Guesses)
}

func TestBadGuess(t *testing.T) {
	game := game.NewGame("blee")
	game, tally := game.Guess('z')

	assert.Equal(t, game.Status(game.Bad), tally.Status)
	assert.Equal(t, 6, tally.TurnsLeft)
	assert.Equal(t, "_ _ _ _", tally.Letters)

	assert.Equal(t, game.Status(game.Bad), game.Status)
	assert.Equal(t, 6, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'z'}, game.Guesses)
}

func TestAlreadyGuessed(t *testing.T) {
	game := game.NewGame("blee")
	game, _ = game.Guess('b')
	game, tally := game.Guess('b')

	assert.Equal(t, game.Status(game.Guessed), tally.Status)
	assert.Equal(t, 7, tally.TurnsLeft)
	assert.Equal(t, "b _ _ _", tally.Letters)

	assert.Equal(t, game.Status(game.Guessed), game.Status)
	assert.Equal(t, 7, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'b'}, game.Guesses)
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

	game := game.NewGame("blee")
	var tally game.Tally
	for _, r := range guesses {
		game, tally = game.Guess(r.g)
		assert.Equal(t, r.s, tally.Status)
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

	game := game.NewGame("blee")
	var tally game.Tally
	for _, r := range guesses {
		game, tally = game.Guess(r.g)
		assert.Equal(t, r.s, tally.Status)
	}
}
