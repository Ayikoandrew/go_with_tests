package selc

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Millisecond):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(a string) time.Duration {
	startTime := time.Now()
	http.Get(a)
	duration := time.Since(startTime)
	return duration
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
