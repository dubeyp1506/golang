package main

import (
	"fmt"
	"time"
)

func LongProcess() {
	for i := range 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	// ========= Timer =========
	// fmt.Println("Timer started")
	// timer := time.NewTimer(5 * time.Second)
	// stopped := timer.Stop()
	// if stopped {
	// 	fmt.Println("Timer is stoped")
	// }
	// timer.Reset(time.Second)
	// fmt.Println(<-timer.C)

	// fmt.Println("Timer expired")


	// ========= Long Process =========
	done := make(chan bool)
	go func() {
		LongProcess()
		done <- true
	}()
	timeout := time.After(3*time.Second)
	select {
	case <-timeout:
		fmt.Println("Timeout")
	case <-done:
		fmt.Println("Long process completed")
	}
}
