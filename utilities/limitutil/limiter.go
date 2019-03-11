package limitutil

import "golang.org/x/time/rate"

func NewRateLimiter(avgTokensPSec int, maxTokensPSec int) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(avgTokensPSec), maxTokensPSec)
}

func WaitForAllow(l *rate.Limiter) {
	for {
		if l.Allow() {
			break
		}
	}
}
