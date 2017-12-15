package dictionary_test

import (
	"testing"

	"github.com/derailed/hangman2/internal/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestCreateGoodFile(t *testing.T) {
	wl, err := dictionary.InitWordList()

	assert.Equal(t, nil, err)
	assert.Equal(t, 8881, len(wl))
}

func TestCreateBadFile(t *testing.T) {
	_, err := dictionary.NewWordList("assets/words_toast.txt")

	assert.EqualError(t, err, "No word list file found!: open assets/words_toast.txt: no such file or directory")
}

func TestRandomWord(t *testing.T) {
	wl, _ := dictionary.InitWordList()

	word := wl.RandomWord()
	if len(word) == 0 {
		t.Fatalf("Unable to fetch random word")
	}
}
