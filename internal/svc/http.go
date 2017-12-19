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
func Call(method, url string, payload io.Reader, res interface{}) error {
	fmt.Println("Calling", url)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	clt := http.DefaultClient
	resp, err := clt.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("doh!! %s failed", url)
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println("RES", res)
	return err
}
