package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// }()

	// // Non-blocking receive
	// select {
	// 	case msg := <- ch:
	// 		fmt.Println("Received", msg)
	// 	default:
	// 		fmt.Println("No message received")


	// Non-blocking send
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
				case d := <- data:
					fmt.Println("Message received", d)
				case <-quit:
					fmt.Println("Quit")
					return
				default:
					fmt.Println("waiting for message...")
					time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true
	time.Sleep(1 * time.Second)
}