package backoff

import (
	"time"
)

// Backoff is an interface that returns the time to timeout after a given number of failed attempts
type Backoff interface {
	// Get gets the backoff duration after the given number of attempts
	Get(attempt uint) time.Duration
}

type exponentialBackoff struct {
	// The start duration to wait after the first Next() call
	Start time.Duration
	// The maximum duration to wait
	Max time.Duration
}

// NewExponentialBackoff returns an exponential backoff object starting with the start duration, doubling the timeout each attempt, up to max.
func NewExponentialBackoff(start time.Duration, max time.Duration) Backoff {
	return &exponentialBackoff{Start: start, Max: max}
}

// Get gets the backoff duration after the given number of attempts
func (b *exponentialBackoff) Get(attempt uint) time.Duration {
	d := b.Start
	for i := uint(0); i < attempt; i++ {
		d *= time.Duration(2)
		if d > b.Max {
			return b.Max
		}
	}
	return d
}
