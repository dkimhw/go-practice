package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil

	/*
		time.After is a very handy function when using select.
		Although it didn't happen in our case you can potentially write code that blocks forever
		if the channels you're listening on never return a value. time. After returns
		a chan (like ping) and will send a signal down it after the amount of time you define.
	*/
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
using struct{} because we don't care about what is returned just a signal that the
we are done with the channel

struct{} here because it's the smallest data type available from a memory perspecive

ALWAYS `make` channel
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
