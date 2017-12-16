package game

import "fmt"

type (
	// Service represent a hangman game interaction
	Service interface {
		DictionaryURL() string
		NewGame(string) (*Game, Tally)
		Guess(*Game, string) (*Game, Tally)
	}

	service struct {
		dicURL string
		game   *Game
	}
)

// NewService creates a game service
func NewService(dicURL string) Service {
	return &service{dicURL: dicURL}
}

func (s *service) DictionaryURL() string {
	return fmt.Sprintf("%s/dictionary/v1/random_word", s.dicURL)
}

// NewGame starts a new hangman game
func (s *service) NewGame(word string) (*Game, Tally) {
	var tally Tally
	s.game, tally = NewGame(word)
	return s.game, tally
}

// Guess a letter
func (s *service) Guess(g *Game, l string) (*Game, Tally) {
	return g.Guess(rune(l[0]))
}
