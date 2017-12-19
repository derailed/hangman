package hangman

import (
	"context"
	"net/http"

	"github.com/derailed/hangman2/internal/game"
	"github.com/derailed/hangman2/internal/svc"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

type (
	// NewGameRequest for a new game
	NewGameRequest struct {
		Cookies []*http.Cookie `json:"cookies"`
	}
	// NewGameResponse returns a new game and tally
	Response struct {
		Game  game.Game `json:"game"`
		Tally Tally     `json:"tally"`
	}
)

func decodeNewGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return NewGameRequest{Cookies: r.Cookies()}, nil
}

func makeNewGameEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewGameRequest)
		ga, tally, err := svc.NewGame(req.Cookies)
		return Response{Game: ga, Tally: tally}, err
	}
}

func makeGuessEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(game.GuessRequest)
		ga, tally, err := svc.Guess(req.Game, req.Guess)
		return Response{Game: ga, Tally: tally}, err
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
		game.DecodeGuessRequest,
		svc.EncodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/game/v1/health", svc.MakeHealthHandler(opts))
	r.Handle("/hangman/v1/new_game", newGameHandler)
	r.Handle("/hangman/v1/guess", guessHandler).Methods("POST")

	return r
}
