package hangman

import (
	"net/http"

	"github.com/derailed/hangman/internal/game"
)

type (
	// Service represent a hangman game interaction
	Service interface {
		NewGame([]*http.Cookie) (game.Game, Tally, error)
		Guess(game.Game, string) (game.Game, Tally, error)
	}

	service struct {
		dicURL, gameURL string
	}
)

// NewService creates a game service
func NewService(dicURL, gameURL string) Service {
	return &service{
		dicURL:  dicURL,
		gameURL: gameURL,
	}
}

func withTally(g game.Game) (game.Game, Tally, error) {
	return g, tallyFromGame(g), nil
}

// NewGame starts a new hangman game
func (s *service) NewGame(cookies []*http.Cookie) (game.Game, Tally, error) {
	word, err := s.NewWord(cookies)
	if err != nil {
		return game.Game{}, Tally{}, err
	}
	g, err := newGame(s.gameURL, word)
	if err != nil {
		return g, Tally{}, err
	}

	return withTally(g)
}

// Guess a letter
func (s *service) Guess(g game.Game, letter string) (game.Game, Tally, error) {
	ng, err := guess(s.gameURL, g, letter)
	if err != nil {
		return ng, Tally{}, err
	}
	return withTally(ng)
}
