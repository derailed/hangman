package game

type (
	// GameService represent a hangman game interaction
	GameService interface {
		NewGame() *Game
		Guess(rune) (*Game, Tally)
	}

	gameService struct {
		list WordList
		game *Game
	}
)

// InitGameSvc for a hangman game
func InitGameSvc() (gameService, error) {
	list, err := InitWordList()
	if err != nil {
		return gameService{}, err
	}
	return gameService{list: list}, nil
}

// NewGame starts a new hangman game
func (s gameService) NewGame() *Game {
	s.game = NewGame(s.list.RandomWord())
	return s.game
}

// Guess a letter
func (s gameService) Guess(l rune) (*Game, Tally) {
	return s.Guess(l)
}
