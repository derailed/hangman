package svc

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

// HealthResponse describes a healthy status
type HealthResponse struct {
	Status string `json:"status"`
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

func makeHealthEP() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return HealthResponse{Status: "ok"}, nil
	}
}

// MakeHealthHandler creates a new status handler
func MakeHealthHandler(opts []kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeHealthEP(),
		decodeHealthRequest,
		EncodeResponse,
		opts...,
	)
}
