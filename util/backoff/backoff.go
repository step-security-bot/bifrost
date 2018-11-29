package backoff

import (
	"time"

	"github.com/cenkalti/backoff"
)

// Construct constructs the backoff.
// Validates the options.
func (b *Backoff) Construct() backoff.BackOff {
	switch b.GetBackoffKind() {
	default:
		fallthrough
	case BackoffKind_BackoffKind_EXPONENTIAL:
		return b.constructExpo()
	case BackoffKind_BackoffKind_CONSTANT:
		return b.constructConstant()
	}
}

// constructExpo constructs an exponential backoff.
func (b *Backoff) constructExpo() backoff.BackOff {
	expo := backoff.NewExponentialBackOff()
	opts := b.GetExponential()

	initialInterval := opts.GetInitialInterval()
	if initialInterval == 0 {
		// default to 500ms
		initialInterval = 500
	}
	expo.InitialInterval = time.Duration(initialInterval) * time.Millisecond

	multiplier := opts.GetMultiplier()
	if multiplier == 0 {
		multiplier = 1.7
	}
	expo.Multiplier = float64(multiplier)

	maxInterval := opts.GetMaxInterval()
	if maxInterval == 0 {
		maxInterval = 10000
	}
	expo.MaxInterval = time.Duration(maxInterval) * time.Millisecond
	expo.RandomizationFactor = float64(opts.GetRandomizationFactor())
	expo.MaxElapsedTime = time.Duration(opts.GetMaxElapsedTime()) * time.Millisecond
	return expo
}

// constructConstant constructs a constant backoff.
func (b *Backoff) constructConstant() backoff.BackOff {
	dur := b.GetConstant().GetInterval()
	if dur == 0 {
		dur = 5000
	}
	return backoff.NewConstantBackOff(time.Duration(dur) * time.Millisecond)
}