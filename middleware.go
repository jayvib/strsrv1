package main

import (
	"time"

	log "github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw *loggingMiddleware) UpperCase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.UpperCase(s)
	return
}

func (mw *loggingMiddleware) Count(s string) (n int) {

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Count(s)
	return
}
