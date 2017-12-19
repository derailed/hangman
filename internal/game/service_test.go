package game_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/derailed/hangman2/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	svc := game.NewService()

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, nil))

	payload := game.NewGameRequest{Word: "fred"}
	req := test.MakeSimpleRequest("POST", MakeURL("game/v1/new_game"), payload)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.NewGameResponse
	recorded.DecodeJsonPayload(&res)
	assert.Equal(t, "fred", res.Game.Letters)
	assert.Equal(t, 7, res.Game.TurnsLeft)
	assert.Equal(t, game.Status(game.Started), res.Game.Status)
	assert.Equal(t, "_ _ _ _", res.Tally.Letters)
}

func TestGuessApi(t *testing.T) {
	svc := game.NewService()

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, nil))

	g := newGame(t, "fred")

	payload := game.GuessRequest{Game: g, Guess: "f"}
	req := test.MakeSimpleRequest("POST", MakeURL("game/v1/guess"), payload)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.GuessResponse
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, 7, res.Tally.TurnsLeft)
	assert.Equal(t, "f _ _ _", res.Tally.Letters)
}

func TestStatusCall(t *testing.T) {
	svc := game.NewService()

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, nil))

	req := test.MakeSimpleRequest("GET", MakeURL("game/v1/health"), nil)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.HealthResponse
	recorded.DecodeJsonPayload(&res)

	assert.Equal(t, "ok", res.Status)
}

func MakeURL(path string) string {
	return fmt.Sprintf("http://%s/%s", httptest.DefaultRemoteAddr, path)
}

func newGame(t *testing.T, word string) game.Game {
	svc := game.NewService()

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, nil))

	payload := game.NewGameRequest{Word: word}
	req := test.MakeSimpleRequest("POST", MakeURL("game/v1/new_game"), payload)
	req.Header.Set("Content-Type", "application/json")

	recorded := test.RunRequest(
		t,
		mux,
		req,
	)

	recorded.CodeIs(http.StatusOK)
	recorded.ContentTypeIsJson()

	var res game.NewGameResponse
	recorded.DecodeJsonPayload(&res)
	return res.Game
}
