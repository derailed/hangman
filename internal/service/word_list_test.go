package service_test

import (
	"testing"

	"github.com/derailed/hangman2/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateGoodFile(t *testing.T) {
	wl, err := service.InitWordList()

	assert.Equal(t, nil, err)
	assert.Equal(t, 8881, len(wl))
}

func TestCreateBadFile(t *testing.T) {
	_, err := service.NewWordList("assets/words_toast.txt")

	assert.EqualError(t, err, "No word list file found!: open assets/words_toast.txt: no such file or directory")
}

func TestRandomWord(t *testing.T) {
	wl, _ := service.InitWordList()

	word := wl.RandomWord()
	if len(word) == 0 {
		t.Fatalf("Unable to fetch random word")
	}
}
