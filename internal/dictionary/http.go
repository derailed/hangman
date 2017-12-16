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

	pingResponse struct {
		Status string `json:"status"`
	}
)

func decodeRandomWordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

func decodePingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeRandomWordEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return randomWordResponse{Word: svc.Word()}, nil
	}
}

func makePingEndPoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return pingResponse{Status: "ok"}, nil
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

	pingHandler := kithttp.NewServer(
		makePingEndPoint(svc),
		decodePingRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/dictionary/v1/health", pingHandler).Methods("GET")
	r.Handle("/dictionary/v1/random_word", randomWordHandler).Methods("GET")

	return r
}
