package retry

import (
	"errors"
	"time"
)

var (
	ErrRetryFailed = errors.New("all retry attempts failed")
)

type Retryer struct {
	totalRetryTimes int
	nextDelay       func() uint32
}

func (r *Retryer) On(method func() error) error {
	retry := 0

	for retry < r.totalRetryTimes {
		err := method()
		if err == nil {
			return nil
		}
		delay := r.nextDelay()
		time.Sleep(time.Duration(delay) * time.Millisecond)

		retry++
	}

	return ErrRetryFailed
}

func NewFixedIntervalRetryer(times int, delay uint32) *Retryer {
	return &Retryer{
		totalRetryTimes: times,
		nextDelay: func() uint32 {
			return delay
		},
	}
}

func NewProgressiveIntervalRetryer(times int, delay uint32) *Retryer {
	next := delay

	return &Retryer{
		totalRetryTimes: times,
		nextDelay: func() uint32 {
			r := next
			next += delay

			return r
		},
	}
}
