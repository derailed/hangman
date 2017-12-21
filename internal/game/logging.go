package game

import (
	"time"

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
func (mw *loggingService) NewGame(word string) Game {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "newGame",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.NewGame(word)
}

// Guess logging wrapper
func (mw *loggingService) Guess(g Game, l rune) Game {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "guess",
			"input", l,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.Guess(g, l)
}
