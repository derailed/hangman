package game

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/derailed/hangman2/internal/svc"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

const dictionary_url = "http://localhost:9090/dictionary/v1/random_word"

type (
	// NewGameRequest requests a new game given a word
	NewGameRequest struct {
		Word    string         `json:"word"`
		Cookies []*http.Cookie `json:"cookies"`
	}
	// Response replies with the current game
	Response struct {
		Game Game `json:"game"`
	}

	// GuessRequest proposes a new guess
	GuessRequest struct {
		Game  Game   `json:"game"`
		Guess string `json:"guess"`
	}
)

func decodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req NewGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.Cookies = r.Cookies()

	return req, nil
}

func DecodeGuessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GuessRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func makeNewGameEndPoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		game := s.NewGame(req.Word)
		return Response{Game: game}, nil
	}
}

func makeGuessEndPoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GuessRequest)
		game := s.Guess(req.Game, rune(req.Guess[0]))
		return Response{Game: game}, nil
	}
}

// MakeHandler to serve game routes
func MakeHandler(s Service, l kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(l),
	}

	newGameHandler := kithttp.NewServer(
		makeNewGameEndPoint(s),
		decodeNewGameRequest,
		svc.EncodeResponse,
		opts...,
	)

	guessHandler := kithttp.NewServer(
		makeGuessEndPoint(s),
		DecodeGuessRequest,
		svc.EncodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/game/v1/health", svc.MakeHealthHandler(opts))
	r.Handle("/game/v1/new_game", newGameHandler).Methods("POST")
	r.Handle("/game/v1/guess", guessHandler).Methods("POST")

	return r
}
