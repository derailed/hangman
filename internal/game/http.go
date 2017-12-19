package game

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

const dictionary_url = "http://localhost:9090/dictionary/v1/random_word"

type (
	NewGameRequest struct {
		Word    string         `json:"word"`
		Cookies []*http.Cookie `json:"cookies"`
	}

	NewGameResponse struct {
		Game  Game  `json:"game"`
		Tally Tally `json:"tally"`
	}

	healthRequest struct {
	}

	HealthResponse struct {
		Status string `json:"status"`
	}

	GuessRequest struct {
		Game  Game   `json:"game"`
		Guess string `json:"guess"`
	}

	GuessResponse struct {
		Game  Game  `json:"game"`
		Tally Tally `json:"tally"`
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

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return healthRequest{}, nil
}

func decodeGuessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GuessRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func makeHealthEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return HealthResponse{Status: "ok"}, nil
	}
}

func makeNewGameEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		game, tally := svc.NewGame(req.Word)
		return NewGameResponse{Game: game, Tally: tally}, nil
	}
}

func makeGuessEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GuessRequest)
		game, tally := svc.Guess(req.Game, req.Guess)
		return GuessResponse{Game: game, Tally: tally}, nil
	}
}

// MakeHandler to serve game routes
func MakeHandler(svc Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	pingHandler := kithttp.NewServer(
		makeHealthEndPoint(svc),
		decodeHealthRequest,
		encodeResponse,
		opts...,
	)

	newGameHandler := kithttp.NewServer(
		makeNewGameEndPoint(svc),
		decodeNewGameRequest,
		encodeResponse,
		opts...,
	)

	guessHandler := kithttp.NewServer(
		makeGuessEndPoint(svc),
		decodeGuessRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/game/v1/health", pingHandler)
	r.Handle("/game/v1/new_game", newGameHandler).Methods("POST")
	r.Handle("/game/v1/guess", guessHandler).Methods("POST")

	return r
}
