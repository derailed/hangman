package dictionary_test

import (
	"testing"

	"github.com/derailed/hangman/internal/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestCreateGoodFile(t *testing.T) {
	wl, err := dictionary.NewWordList("assets/words.txt")

	assert.Equal(t, nil, err)
	assert.Equal(t, 5, len(wl))
}

func TestCreateBadFile(t *testing.T) {
	_, err := dictionary.NewWordList("assets/words_toast.txt")

	assert.EqualError(t, err, "No word list file found!: open assets/words_toast.txt: no such file or directory")
}

func TestRandomWord(t *testing.T) {
	wl, _ := dictionary.NewWordList("assets/words.txt")

	word := wl.Word()
	if len(word) == 0 {
		t.Fatalf("Unable to fetch random word")
	}
}
