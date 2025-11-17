package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func increment(counter *int32) {
	atomic.AddInt32(counter, 1)
}

func decrement(counter *int32) {
	atomic.AddInt32(counter, -1)
}

func getValue(counter *int32) int32 {
	return atomic.LoadInt32(counter)
}

func main() {
	var counter int32
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				increment(&counter)
			}
			decrement(&counter)
		}()

	}

	wg.Wait()
	fmt.Println(getValue(&counter))
}
