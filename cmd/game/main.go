package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/derailed/hangman/internal/cors"
	"github.com/derailed/hangman/internal/game"
	"github.com/go-kit/kit/log"
)

const port = ":9095"

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc := game.NewService()
	svc = game.NewLoggingService(svc, logger)

	mux := http.NewServeMux()
	mux.Handle("/game/v1/", game.MakeHandler(svc, logger))

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
