package main

import (
	"fmt"
	"time"
)
type Ticket struct {
	id int
	price int
}

// ========= Ticket processing using Worker Pools =========

func ticketProcessor(request <- chan Ticket, results chan <- int) {
	for req := range request {
		fmt.Printf("Processing ticket %d with price %d\n", req.id, req.price)
		time.Sleep(time.Second)
		results <- req.id
	}
}

func main() {
	const numWorkers = 4
	const numTasks = 10

	request := make(chan Ticket, numTasks)
	results := make(chan int, numTasks)

	for i := 1; i <= numWorkers; i++ {
		go ticketProcessor(request, results)
	}

	for i := 1; i <= numTasks; i++ {
		request <- Ticket{id: i, price: i*10}
	}
	close(request)

	for i := 1; i <= numTasks; i++ {
		fmt.Println(<-results)
	}
}

// ========= Basic Worker Pools =========

// func worker(id int , tasks <- chan int, results chan <- int) {
// 	task := <-tasks
// 	fmt.Printf("Worker %d started with task %d\n", id, task)
// 	time.Sleep(2*time.Second)
// 	results <- task * 2
// }

// func main() {
// 	const numWorkers = 2
// 	const numTasks = 10

// 	tasks := make(chan int, numTasks)
// 	results := make(chan int, numTasks)

// 	for i := 1; i <= numTasks; i++ {
// 		go worker(i, tasks, results)
// 	}

// 	for i := 1; i <= numTasks; i++ {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	for i:=1; i<=numTasks; i++ {
// 		fmt.Println(<-results)
// 	}
// }
