package dictionary

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	Randomizer
	logger log.Logger
}

// NewLoggingService returns a new instance to the logging service
func NewLoggingService(s Randomizer, l log.Logger) Randomizer {
	return &loggingService{s, l}
}

// NewGame logging wrapper
func (mw *loggingService) Word() (word string) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "word",
			"word", word,
			"took", time.Since(begin),
		)
	}(time.Now())

	word = mw.Randomizer.Word()
	return
}
