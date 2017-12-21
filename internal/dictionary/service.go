package dictionary

import "github.com/pkg/errors"

type service struct {
	list WordList
}

// NewService creates a new dictionary service
func NewService(path string) (Randomizer, error) {
	list, err := NewWordList(path)
	if err != nil {
		return &service{}, errors.Wrap(err, "no word list file found!")
	}
	return &service{list: list}, nil
}

// NewWord fetches a new random word for the words list
func (s *service) Word() string {
	return s.list.Word()
}
