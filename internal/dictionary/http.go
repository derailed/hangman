package dictionary

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type (
	randomWordResponse struct {
		Word string `json:"word"`
	}
)

// DecodeRandomWordRequest - no opt
func DecodeRandomWordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return "", nil
}

// EncodeResponse to json
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// MakeRandomWordEndPoint create a endpoint for a service
func MakeRandomWordEndPoint(svc RandomWordService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return randomWordResponse{Word: svc.Word()}, nil
	}
}
