package hangman

import (
	"net/http"
	"time"

	"github.com/derailed/hangman/internal/game"
	"github.com/go-kit/kit/log"
)

type loggingService struct {
	Service
	logger log.Logger
}

// NewLoggingService returns a new instance to the logging service
func NewLoggingService(s Service, l log.Logger) Service {
	return &loggingService{s, l}
}

// NewGame logging wrapper
func (mw *loggingService) NewGame(cookies []*http.Cookie) (game.Game, Tally, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "new_game",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.NewGame(cookies)
}

// Guess logging wrapper
func (mw *loggingService) Guess(g game.Game, l string) (game.Game, Tally, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "guess",
			"input", l,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.Guess(g, l)
}
