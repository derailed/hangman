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
	_, err := dictionary.NewService("assets/word.txt")
	if err == nil {
		t.Fatalf("Should flag no assets")
	}
}

func TestRandomWordCall(t *testing.T) {
	mux := createMux(t)

	recorded := test.RunRequest(t, mux, makeRequest("GET", "random_word"))
	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	res := map[string]string{}
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, "this", res["word"])
}

func TestStatusCall(t *testing.T) {
	mux := createMux(t)

	recorded := test.RunRequest(t, mux, makeRequest("GET", "health"))
	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	res := map[string]string{}
	recorded.DecodeJsonPayload(&res)
	assert.Equal(t, "ok", res["status"])
}

func MakeURL(action string) string {
	return fmt.Sprintf("http://%s/dictionary/v1/%s", httptest.DefaultRemoteAddr, action)
}

func makeRequest(method, action string) *http.Request {
	req := test.MakeSimpleRequest(method, MakeURL(action), nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func createMux(t *testing.T) *http.ServeMux {
	svc, err := dictionary.NewService("assets/words.txt")
	assert.Equal(t, nil, err)

	mux := http.NewServeMux()
	mux.Handle("/dictionary/v1/", dictionary.MakeHandler(svc, nil))

	return mux
}
