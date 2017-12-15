package dictionary

type (
	// RandomWordService interact with an underlying word list
	RandomWordService interface {
		Word() string
	}

	randomWordService struct {
		list WordList
	}
)

func InitSvc() (randomWordService, error) {
	list, err := InitWordList()
	if err != nil {
		return randomWordService{}, err
	}
	return randomWordService{list: list}, nil
}

// NewWord fetches a new random word for the words list
func (s randomWordService) Word() string {
	return s.list.RandomWord()
}
