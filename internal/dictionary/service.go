package dictionary

type (
	// RandomWordService interact with an underlying word list
	Service interface {
		Word() string
	}

	service struct {
		list WordList
	}
)

// New dictionary service
func New() (Service, error) {
	list, err := InitWordList()
	if err != nil {
		return service{}, err
	}
	return service{list: list}, nil
}

// NewWord fetches a new random word for the words list
func (s service) Word() string {
	return s.list.RandomWord()
}
