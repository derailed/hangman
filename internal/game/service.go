package game

type (
	// Service represent a hangman game interaction
	Service interface {
		NewGame(string) *Game
		Guess(*Game, string) (*Game, Tally)
	}

	service struct {
		game *Game
	}
)

// NewService creates a game service
func NewService() Service {
	return &service{}
}

// NewGame starts a new hangman game
func (s *service) NewGame(word string) *Game {
	s.game = NewGame(word)
	return s.game
}

// Guess a letter
func (s *service) Guess(g *Game, l string) (*Game, Tally) {
	return g.Guess(rune(l[0]))
}
