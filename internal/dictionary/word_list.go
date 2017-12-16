package dictionary

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

const path = "assets/words.txt"

type (
	// Randomizer selects a word at random
	Randomizer interface {
		randomWord() string
	}
	// WordList stores a collection of words
	WordList []string
)

// Init dictionary from default word list
func InitWordList() (WordList, error) {
	return NewWordList(path)
}

// NewWordList creates a new wordlist from the given file
func NewWordList(path string) (WordList, error) {
	wl, err := load(path)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

// RandomWord returns a new word randomly from the list
func (wl WordList) RandomWord() string {
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
