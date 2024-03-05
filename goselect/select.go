package goselect

import (
	"errors"
	"net/http"
	"time"
)

var TimeOutError = errors.New("error timed out in waiting for response from the urls")

const timeout = 10 * time.Second

func Racer(a, b string) string {
	aDuration := measureTime(a)
	bDuration := measureTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

// select will be used when we are waiting on multiple channels
func RacerWithSelect(a, b string) (string, error) {
	return ConfigurableRacer(a, b, timeout)
}

// we didn't needto care timeout in happy case so we will defined a timeout const and will use that
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", TimeOutError
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(url string) {
		http.Get(url)
		close(ch)
	}(url)

	return ch
}
