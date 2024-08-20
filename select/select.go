package main

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}
