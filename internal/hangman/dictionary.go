package hangman

import (
	"fmt"
	"net/http"

	"github.com/derailed/hangman/internal/dictionary"
	"github.com/derailed/hangman/internal/svc"
)

func dicURL(url, path string) string {
	return fmt.Sprintf("%s/dictionary/v1/%s", url, path)
}

// NewWord retrieves a random word from dictionary service
func (s *service) NewWord(cookies []*http.Cookie) (string, error) {
	var res dictionary.WordResponse
	err := svc.Call("GET", dicURL(s.dicURL, "random_word"), nil, &res)
	return res.Word, err
}
