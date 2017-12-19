package hangman

import (
	"net/http"
	"time"

	"github.com/derailed/hangman2/internal/game"
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
func (mw loggingService) NewGame(cookies []*http.Cookie) (int, game.Tally, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "new_game",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.NewGame(cookies)
}

// Guess logging wrapper
func (mw loggingService) Guess(id string, letter string) (game.Tally, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "guess",
			"game_id", id,
			"input", letter,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.Guess(id, letter)
}

func (mw loggingService) EndGame(id string) error {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "end_game",
			"game_id", id,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Service.EndGame(id)
}
