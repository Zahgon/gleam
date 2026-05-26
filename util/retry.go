package util

import (
	"time"
)

func Retry(fn func() error) error { _ = "STUB: not implemented"; return nil }

func TimeDelayedRetry(fn func() error, waitTimes ...time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
