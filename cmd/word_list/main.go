package main

import (
	"net/http"
	"os"

	"github.com/derailed/hangman2/internal/service"
	"github.com/derailed/hangman2/internal/transport"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

const port = ":9090"

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.RandomWordService

	svc, err := service.InitSvc()
	if err != nil {
		panic(err)
	}

	randomWordHandler := kithttp.NewServer(
		transport.MakeRandomWordEndPoint(svc),
		transport.DecodeRandomWordRequest,
		transport.EncodeResponse,
	)

	http.Handle("/random_word", randomWordHandler)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
