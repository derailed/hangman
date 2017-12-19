package dictionary

import (
	"context"
	"net/http"

	"github.com/derailed/hangman/internal/svc"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// WordResponse returns a new word
type WordResponse struct {
	Word string `json:"word"`
}

func decodeRandomWordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

func makeRandomWordEndPoint(svc Randomizer) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return WordResponse{Word: svc.Word()}, nil
	}
}

// MakeHandler to service route requests
func MakeHandler(s Randomizer, l kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(l),
	}

	randomWordHandler := kithttp.NewServer(
		makeRandomWordEndPoint(s),
		decodeRandomWordRequest,
		svc.EncodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/dictionary/v1/health", svc.MakeHealthHandler(opts)).Methods("GET")
	r.Handle("/dictionary/v1/random_word", randomWordHandler).Methods("GET")

	return r
}
