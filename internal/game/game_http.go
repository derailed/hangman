package game

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type (
	newGameResponse struct {
		Game *Game `json:"game"`
	}

	guessRequest struct {
		Guess string `json:"guess"`
	}

	guessResponse struct {
		Game  *Game `json:"game"`
		Tally Tally `json:"tally"`
	}
)

// DecodeNewGameRequest - no opt
func DecodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

// DecodeGuessRequest - no opt
func DecodeGuessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req guessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// EncodeResponse to json
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeNewGameEndPoint create new game endpoint for a service
func MakeNewGameEndPoint(svc GameService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return newGameResponse{Game: svc.NewGame()}, nil
	}
}

// MakeGuessEndPoint creates a guess endpoint for a service
func MakeGuessEndPoint(svc GameService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(guessRequest)
		game, tally := svc.Guess(req.Guess)
		return guessResponse{Game: game, Tally: tally}, nil
	}
}
