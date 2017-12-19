package dictionary

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

type (
	// Randomizer selects a random word for a collection of words
	Randomizer interface {
		Word() string
	}
	// WordList stores a collection of words
	WordList []string
)

// NewWordList creates a new wordlist from the given file
func NewWordList(path string) (WordList, error) {
	wl, err := load(path)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

// Word returns a new word randomly from the list
func (wl WordList) Word() string {
	return wl[rand.Intn(len(wl))]
}

func load(path string) (WordList, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "No word list file found!")
	}

	var wl WordList
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wl = append(wl, sc.Text())
	}

	return wl, nil
}
