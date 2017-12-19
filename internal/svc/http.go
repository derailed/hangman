package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EncodeResponse returns a json reply
func EncodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// Call a web service with the given url and parameters
func Call(method, url string, payload io.Reader, res interface{}, cookie []*http.Cookie) error {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	if cookie != nil {
		for _, c := range cookie {
			req.AddCookie(c)
		}
	}

	clt := http.DefaultClient
	resp, err := clt.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("doh!! %s failed", url)
	}

	return json.NewDecoder(resp.Body).Decode(&res)
}
