package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/derailed/hangman2/internal/cors"
	"github.com/derailed/hangman2/internal/dictionary"
	"github.com/go-kit/kit/log"
)

const port = ":9090"

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	svc, err := dictionary.New()
	if err != nil {
		panic(err)
	}
	svc = dictionary.NewLoggingService(svc, logger)

	mux := http.NewServeMux()
	mux.Handle("/dictionary/v1/", dictionary.MakeHandler(svc, logger))
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
