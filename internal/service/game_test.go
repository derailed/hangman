package service_test

import (
	"testing"

	"github.com/derailed/hangman2/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGoodGuess(t *testing.T) {
	game := service.NewGame("blee")
	game, tally := game.Guess('e')

	assert.Equal(t, service.Status(service.Good), tally.Status)
	assert.Equal(t, 7, tally.TurnsLeft)
	assert.Equal(t, "_ _ e e", tally.Letters)

	assert.Equal(t, service.Status(service.Good), game.Status)
	assert.Equal(t, 7, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'e'}, game.Guesses)
}

func TestBadGuess(t *testing.T) {
	game := service.NewGame("blee")
	game, tally := game.Guess('z')

	assert.Equal(t, service.Status(service.Bad), tally.Status)
	assert.Equal(t, 6, tally.TurnsLeft)
	assert.Equal(t, "_ _ _ _", tally.Letters)

	assert.Equal(t, service.Status(service.Bad), game.Status)
	assert.Equal(t, 6, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'z'}, game.Guesses)
}

func TestAlreadyGuessed(t *testing.T) {
	game := service.NewGame("blee")
	game, _ = game.Guess('b')
	game, tally := game.Guess('b')

	assert.Equal(t, service.Status(service.Guessed), tally.Status)
	assert.Equal(t, 7, tally.TurnsLeft)
	assert.Equal(t, "b _ _ _", tally.Letters)

	assert.Equal(t, service.Status(service.Guessed), game.Status)
	assert.Equal(t, 7, game.TurnsLeft)
	assert.Equal(t, "blee", game.Letters)
	assert.Equal(t, []rune{'b'}, game.Guesses)
}

func TestWin(t *testing.T) {
	guesses := []struct {
		g rune
		s service.Status
	}{
		{g: 'b', s: service.Good},
		{g: 'l', s: service.Good},
		{g: 'e', s: service.Won},
	}

	game := service.NewGame("blee")
	var tally service.Tally
	for _, r := range guesses {
		game, tally = game.Guess(r.g)
		assert.Equal(t, r.s, tally.Status)
	}
}
func TestLoose(t *testing.T) {
	guesses := []struct {
		g rune
		s service.Status
	}{
		{g: 'q', s: service.Bad},
		{g: 'r', s: service.Bad},
		{g: 's', s: service.Bad},
		{g: 'u', s: service.Bad},
		{g: 'v', s: service.Bad},
		{g: 'x', s: service.Bad},
		{g: 'z', s: service.Lost},
	}

	game := service.NewGame("blee")
	var tally service.Tally
	for _, r := range guesses {
		game, tally = game.Guess(r.g)
		assert.Equal(t, r.s, tally.Status)
	}
}
