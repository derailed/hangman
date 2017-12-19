package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/derailed/hangman/internal/cors"
	"github.com/derailed/hangman/internal/hangman"
	"github.com/go-kit/kit/log"
)

const port = ":9096"

const localDic = "http://localhost:9094"
const localGame = "http://localhost:9095"

func main() {
	dicURL := flag.String("dicUrl", localDic, "Dictionary Host URL")
	gameURL := flag.String("gameUrl", localGame, "Game Host URL")

	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc := hangman.NewService(*dicURL, *gameURL)
	svc = hangman.NewLoggingService(svc, logger)

	mux := http.NewServeMux()
	mux.Handle("/hangman/v1/", hangman.MakeHandler(svc, logger))

	// Register pprof handlers
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.Handle("/", cors.AccessControl(mux))

	errs := make(chan error, 2)
	go func() {
		logger.Log("msg", "HTTP", "addr", port)
		errs <- http.ListenAndServe(port, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
