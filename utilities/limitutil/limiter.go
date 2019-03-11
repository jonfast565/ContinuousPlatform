package limitutil

import "golang.org/x/time/rate"

// Creates a new rate limiter
func NewRateLimiter(avgTokensPSec int, maxTokensPSec int) *rate.Limiter {
	return rate.NewLimiter(rate.Limit(avgTokensPSec), maxTokensPSec)
}

// Waits for the rate limiter to allow the passage of a method or goroutine
func WaitForAllow(l *rate.Limiter) {
	for {
		if l.Allow() {
			break
		}
	}
}
