package game

type (
	service struct{}

	// Service represents the game service
	Service interface {
		NewGame(string) Game
		Guess(Game, rune) Game
	}
)

// NewService creates a game service
func NewService() Service {
	return &service{}
}

// NewGame starts a new hangman game
func (s *service) NewGame(word string) Game {
	return *NewGame(word)
}

// Guess the next letter
func (s *service) Guess(g Game, l rune) Game {
	return *g.Guess(l)
}
