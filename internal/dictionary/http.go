package dictionary

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

type (
	randomWordResponse struct {
		Word string `json:"word"`
	}
)

// DecodeRandomWordRequest - no opt
func decodeRandomWordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

// EncodeResponse to json
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeRandomWordEndPoint create a endpoint for a service
func makeRandomWordEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return randomWordResponse{Word: svc.Word()}, nil
	}
}

// MakeHandler to service route requests
func MakeHandler(svc Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	randomWordHandler := kithttp.NewServer(
		makeRandomWordEndPoint(svc),
		decodeRandomWordRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/dictionary/v1/random_word", randomWordHandler).Methods("GET")

	return r
}
