package game

type (
	// Service represent a hangman game interaction
	Service interface {
		NewGame(string) (Game, Tally)
		Guess(Game, string) (Game, Tally)
	}

	service struct {
	}
)

// NewService creates a game service
func NewService() Service {
	return &service{}
}

// NewGame starts a new hangman game
func (s *service) NewGame(word string) (Game, Tally) {
	game, tally := NewGame(word)
	return game, tally
}

// Guess a letter
func (s *service) Guess(g Game, l string) (Game, Tally) {
	return g.Guess(rune(l[0]))
}
