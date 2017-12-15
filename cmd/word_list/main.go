package main

import (
	"net/http"
	"os"

	"github.com/derailed/hangman2/internal/dictionary"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

const port = ":9090"

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc dictionary.RandomWordService

	svc, err := dictionary.InitSvc()
	if err != nil {
		panic(err)
	}

	randomWordHandler := kithttp.NewServer(
		dictionary.MakeRandomWordEndPoint(svc),
		dictionary.DecodeRandomWordRequest,
		dictionary.EncodeResponse,
	)

	http.Handle("/random_word", randomWordHandler)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
