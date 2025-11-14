package main

import (
	"fmt"
)

func send(ch chan <-int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int , out chan <-int) {
	for i := range in {
		if i % 2 == 0 {
			out <- i
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go send(ch1)
	go filter(ch1, ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	//1 basic closing
	// ch := make(chan int)
	
	// go func() {
	// 	for i := range 5 {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// for i := range ch {
	// 	fmt.Println(i)
	// }
}