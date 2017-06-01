package fcm

import (
	"net"
	"time"
)

const (
	minBackoff = 100 * time.Millisecond
	maxBackoff = 1 * time.Minute
	factor     = 2.7
)

func retry(fn func() error, attempts int) error {
	var attempt int
	for {
		err := fn()
		if err == nil {
			return nil
		}

		if tErr, ok := err.(net.Error); !ok || !tErr.Temporary() {
			return err
		}

		attempt++
		backoff := minBackoff * time.Duration(attempt*attempt)
		if attempt > attempts || backoff > maxBackoff {
			return err
		}

		time.Sleep(backoff)
	}
}
