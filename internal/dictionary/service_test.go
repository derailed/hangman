package dictionary_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/derailed/hangman2/internal/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestNoAssets(t *testing.T) {
	_, err := dictionary.New("assets/word.txt")
	if err == nil {
		t.Fatalf("Should flag no assets")
	}
}

func TestRandomWordCall(t *testing.T) {
	svc, err := dictionary.New("assets/words.txt")
	assert.Equal(t, nil, err)

	mux := http.NewServeMux()
	mux.Handle("/dictionary/v1/", dictionary.MakeHandler(svc, nil))

	req := test.MakeSimpleRequest("GET", MakeURL("dictionary/v1/random_word"), nil)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	res := map[string]string{}
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, "this", res["word"])
}

func TestStatusCall(t *testing.T) {
	svc, err := dictionary.New("assets/words.txt")
	assert.Equal(t, nil, err)

	mux := http.NewServeMux()
	mux.Handle("/dictionary/v1/", dictionary.MakeHandler(svc, nil))

	req := test.MakeSimpleRequest("GET", MakeURL("dictionary/v1/health"), nil)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	res := map[string]string{}
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, "ok", res["status"])
}

func MakeURL(path string) string {
	return fmt.Sprintf("http://%s/%s", httptest.DefaultRemoteAddr, path)
}
