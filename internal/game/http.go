package game

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

const dictionary_url = "http://localhost:9090/dictionary/v1/random_word"

type (
	newGameRequest struct {
		Word string `json:"word"`
	}

	newGameResponse struct {
		Game *Game `json:"game"`
	}

	guessRequest struct {
		Game  *Game  `json:"game"`
		Guess string `json:"guess"`
	}

	guessResponse struct {
		Game  *Game `json:"game"`
		Tally Tally `json:"tally"`
	}
)

func decodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGuessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req guessRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeNewGameEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		word, err := getWord("http://dictionary:9090/dictionary/v1/random_word")
		if err != nil {
			return nil, err
		}
		return newGameResponse{Game: svc.NewGame(word)}, nil
	}
}

func getWord(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	clt := http.DefaultClient
	resp, err := clt.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("dictionary toast")
	}

	res := map[string]interface{}{}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res["word"].(string), nil
}

func makeGuessEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(guessRequest)
		game, tally := svc.Guess(req.Game, req.Guess)
		return guessResponse{Game: game, Tally: tally}, nil
	}
}

// MakeHandler to serve game routes
func MakeHandler(svc Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

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

	r.Handle("/game/v1/new_game", newGameHandler).Methods("GET")
	r.Handle("/game/v1/guess", guessHandler).Methods("POST")

	return r
}
