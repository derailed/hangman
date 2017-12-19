package game_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/derailed/hangman/internal/game"
	"github.com/derailed/hangman/internal/svc"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := newGame(t, "fred")

	assert.Equal(t, "fred", g.Letters)
	assert.Equal(t, 7, g.TurnsLeft)
	assert.Equal(t, game.Status(game.Started), g.Status)
}

func TestGuessApi(t *testing.T) {
	mux := makeMux()
	g := newGame(t, "fred")
	payload := game.GuessRequest{Game: g, Guess: "f"}
	recorded := test.RunRequest(t, mux, makeRequest("POST", "guess", payload))
	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.Response
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, 7, res.Game.TurnsLeft)
	assert.Equal(t, "fred", res.Game.Letters)
}

func TestStatusCall(t *testing.T) {
	mux := makeMux()

	recorded := test.RunRequest(t, mux, makeRequest("GET", "health", nil))
	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res svc.HealthResponse
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, "ok", res.Status)
}

func newGame(t *testing.T, word string) game.Game {
	mux := makeMux()

	payload := game.NewGameRequest{Word: "fred"}
	recorded := test.RunRequest(t, mux, makeRequest("POST", "new_game", payload))
	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.Response
	recorded.DecodeJsonPayload(&res)
	return res.Game
}

func makeURL(action string) string {
	return fmt.Sprintf("http://%s/game/v1/%s", httptest.DefaultRemoteAddr, action)
}

func makeMux() *http.ServeMux {
	svc := game.NewService()

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, nil))
	return mux
}

func makeRequest(method, action string, payload interface{}) *http.Request {
	req := test.MakeSimpleRequest(method, makeURL(action), payload)
	req.Header.Set("Content-Type", "application/json")
	return req
}
