// I am not overly concerned about the vulnerability of the random used in the backoff algorithm.
// nolint: gosec
package jitter

import (
	"math/rand"
	"time"
)

// Delay returns a random duration between (delay * (1 - factor), delay].
func Delay(delay time.Duration, factor float64) time.Duration {
	maxJitter := time.Duration(float64(delay) * factor)

	if maxJitter == 0 {
		return delay
	}

	return delay - UpTo(maxJitter) + 1
}

// UpTo returns a random duration between (0, max].
func UpTo(max time.Duration) time.Duration {
	return time.Duration(rand.Int63n(int64(max))) + 1
}
