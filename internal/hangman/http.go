package hangman

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/derailed/hangman2/internal/game"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

type (
	NewGameRequest struct {
		Cookies []*http.Cookie `json:"cookies"`
	}
	NewGameResponse struct {
		ID    int        `json:"id"`
		Tally game.Tally `json:"tally"`
	}

	EndGameRequest struct {
		ID string `json:"id"`
	}
	EndGameResponse struct {
	}

	pingRequest struct {
	}
	pingResponse struct {
		Status string `json:"status"`
	}

	GuessRequest struct {
		ID    string `json:"id"`
		Guess string `json:"guess"`
	}
	GuessResponse struct {
		Tally game.Tally `json:"tally"`
	}
)

func decodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return NewGameRequest{Cookies: r.Cookies()}, nil
}

func decodeEndGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req EndGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodePingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return pingRequest{}, nil
}

func decodeGuessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GuessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("CRAP", err)
		return nil, err
	}
	fmt.Println("REQ", req)
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makePingEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return pingResponse{Status: "ok"}, nil
	}
}

func makeNewGameEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		id, tally, err := svc.NewGame(req.Cookies)
		return NewGameResponse{ID: id, Tally: tally}, err
	}
}

func makeEndGameEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EndGameRequest)
		err := svc.EndGame(req.ID)
		return EndGameResponse{}, err
	}
}

func makeGuessEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GuessRequest)
		tally, err := svc.Guess(req.ID, req.Guess)
		return GuessResponse{Tally: tally}, err
	}
}

// MakeHandler to serve game routes
func MakeHandler(svc Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	pingHandler := kithttp.NewServer(
		makePingEndPoint(svc),
		decodePingRequest,
		encodeResponse,
		opts...,
	)

	newGameHandler := kithttp.NewServer(
		makeNewGameEndPoint(svc),
		decodeNewGameRequest,
		encodeResponse,
		opts...,
	)

	endGameHandler := kithttp.NewServer(
		makeEndGameEndPoint(svc),
		decodeEndGameRequest,
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

	r.Handle("/hangman/v1/health", pingHandler)
	r.Handle("/hangman/v1/new_game", newGameHandler)
	r.Handle("/hangman/v1/end_game", endGameHandler).Methods("POST")
	r.Handle("/hangman/v1/guess", guessHandler).Methods("POST")

	return r
}
