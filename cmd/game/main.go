package main

import (
	"net/http"
	"os"

	"github.com/derailed/hangman2/internal/game"
	"github.com/go-kit/kit/log"
)

const port = ":9095"

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc game.GameService

	svc, err := game.InitGameSvc()
	if err != nil {
		panic(err)
	}

	newGameHandler := kithttp.NewServer(
		game.MakeNewGameEndPoint(svc),
		game.DecodeNewGameRequest,
		game.EncodeResponse,
	)

	guessHandler := kithttp.NewServer(
		game.MakeGuessEndPoint(svc),
		game.DecodeGuessRequest,
		game.EncodeResponse,
	)

	http.Handle("/new_game", newGameHandler)
	http.Handle("/guess", guessHandler)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
