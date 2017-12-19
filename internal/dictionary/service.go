package dictionary

type (
	// Service interact with an underlying word list
	Service interface {
		Word() string
	}

	service struct {
		list WordList
	}
)

// NewService creates a new dictionary service
func NewService(path string) (Service, error) {
	list, err := NewWordList(path)
	if err != nil {
		return service{}, err
	}
	return service{list: list}, nil
}

// NewWord fetches a new random word for the words list
func (s service) Word() string {
	return s.list.RandomWord()
}
