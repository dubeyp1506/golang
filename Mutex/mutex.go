package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}
func main() {
	counter := &Counter{val: 0}
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				counter.increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.getValue())
}
