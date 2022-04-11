package jitter

import (
	"math/rand"
	"time"
)

// Delay returns a random duration between (delay * (1 - factor), delay].
func Delay(delay time.Duration, factor float64) time.Duration {
	if factor <= 0.0 {
		return delay
	}

	return By(delay, time.Duration(float64(delay)*factor))
}

// By returns a random duration between (delay-jitter, delay].
func By(delay time.Duration, maxJitter time.Duration) time.Duration {
	if maxJitter <= 0 {
		return delay
	}

	if maxJitter > delay {
		return UpTo(delay)
	}

	return delay - UpTo(maxJitter) + 1
}

// UpTo returns a random duration between (0, max].
func UpTo(max time.Duration) time.Duration {
	if max <= 0 {
		return 1
	}

	return time.Duration(rand.Int63n(int64(max))) + 1
}
