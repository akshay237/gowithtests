package gosync

import "sync"

// Sync package is used to access the resources safely b/w multiple goroutines.
/*
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
*/

// When to use Channels and mutexs
/*
- Use channels when passing ownership of data
- Use mutexes for managing state
*/

// Use go vet in your build scripts as it can alert you to some subtle bugs in your code

type Counter struct {
	mutex sync.Mutex
	item  int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.item++
}

func (c *Counter) Value() int {
	return c.item
}

func NewCounter() *Counter {
	return &Counter{}
}
