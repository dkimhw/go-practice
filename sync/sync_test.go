package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// A WaitGroup waits for a collection of goroutines to finish.
		var wg sync.WaitGroup
		// The main goroutine calls Add to set the number of goroutines to wait for.
		wg.Add(wantedCount) // wait for 1000 go routines

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		// By waiting for wg.Wait() to finish before making our assertions we
		// can be sure all of our goroutines have attempted to Inc the Counter.
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
