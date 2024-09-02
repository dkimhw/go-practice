package sync

import "sync"

// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// acquire lock on counter
	c.mu.Lock()
	// all other go routines will have to wait for it to be unlocked before getting access
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
