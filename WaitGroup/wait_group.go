package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id   int
	Task string
}

func (w *Worker) DoWork(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started with task %s\n", w.id, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished with task %s\n", w.id, w.Task)
}

func withChannel(id int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker with channel %d started\n", id)
	time.Sleep(time.Second)
	result <- id
}

// ======== Basic WaitGroup ========
func worker(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker without channel %d started\n", val)
	time.Sleep(time.Second)
}

func main() {
	// ======== Basic WaitGroup ========
	var wg sync.WaitGroup
	fmt.Println("starting workers")
	for i := range 5 {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Worker without channel finished")

	// ======== WaitGroup with Channel ========

	result := make(chan int, 5)
	for i := range 5 {
		wg.Add(1)
		go withChannel(i, result, &wg)
	}

	wg.Wait()
	close(result)

	for v := range result {
		fmt.Println("result received", v)
	}
	fmt.Println("Worker with channel finished")

	// ======== Worker with WaitGroup ========

	tasks := []Worker{
		{id: 1, Task: "task 1"},
		{id: 2, Task: "task 2"},
		{id: 3, Task: "task 3"},
		{id: 4, Task: "task 4"},
		{id: 5, Task: "task 5"},
	}

	for i := range tasks {
		wg.Add(1)
		go tasks[i].DoWork(&wg)
	}
	wg.Wait()
	fmt.Println("Worker tasks finished")
}
