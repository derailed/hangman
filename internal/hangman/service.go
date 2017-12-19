package hangman

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/derailed/hangman2/internal/game"
)

type (
	// Service represent a hangman game interaction
	Service interface {
		NewGame([]*http.Cookie) (int, game.Tally, error)
		Guess(string, string) (game.Tally, error)
		EndGame(string) error
	}

	service struct {
		dicURL, gameURL string
		id              int
		games           map[int]game.Game
	}
)

// NewService creates a game service
func NewService(dicURL, gameURL string) Service {
	return &service{
		dicURL:  dicURL,
		gameURL: gameURL,
		games:   make(map[int]game.Game),
	}
}

// NewGame starts a new hangman game
func (s *service) NewGame(cookies []*http.Cookie) (int, game.Tally, error) {
	word, err := s.NewWord(cookies)
	if err != nil {
		return 0, game.Tally{}, err
	}
	fmt.Println("WORD", word)
	g, tally, err := newGame(s.gameURL, word)
	if err != nil {
		return 0, game.Tally{}, err
	}
	id := s.nextID()
	s.games[id] = g

	return id, tally, nil
}

// Guess a letter
func (s *service) Guess(id string, letter string) (game.Tally, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return game.Tally{}, err
	}
	g, ok := s.games[index]
	if !ok {
		return game.Tally{}, fmt.Errorf("Unable to find game id `%d", index)
	}
	game, tally, err := guess(s.gameURL, g, letter)
	if err != nil {
		return tally, err
	}
	s.games[index] = game
	return tally, err
}

func (s *service) EndGame(id string) error {
	index, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	delete(s.games, index)
	return nil
}

func (s *service) nextID() int {
	s.id++
	return s.id
}
