package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func isRunning(ctx context.Context) {
	for {
	
		select {
		case <-ctx.Done():
			fmt.Println("Job is cancelled", ctx.Err())
			return
		default:
			fmt.Println("Job is Running")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func isEven(ctx context.Context, num int) string {
	select {
		case  <-ctx.Done():
			fmt.Println("ctx.Err():", ctx.Err())
			return "Context Done"
		default:
		if num&1 == 0 {
			return fmt.Sprintf("%d is even", num)
		}else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func logger(ctx context.Context, message string) {
	reqValue := ctx.Value("request_id")
	log.Println(ctx, reqValue, message)
}

func main() {

	//========= WithTimeOut ==========
	// ctx := context.TODO()

	// result := isEven(ctx, 10)
	// fmt.Println(result)

	// ctx = context.Background()
	// ctx, cancel := context.WithTimeout(ctx,1*time.Second)
	// fmt.Println(cancel)
	// defer cancel()

	// time.Sleep(2*time.Second)
	// result = isEven(ctx, 11)
	// fmt.Println(result)

	//========== Background context ==========

	// ctx := context.Background()
	// ctx, cancle	:= context.WithTimeout(ctx,1*time.Second)
	// defer cancle()

	// go isRunning(ctx)
	// ctx = context.WithValue(ctx, "user", "Priyanshu")
	// time.Sleep(2000* time.Millisecond)

	// if ctx.Value("user") != nil {
	// 	fmt.Println(ctx.Value("user"))
	// 	return
	// }
	// fmt.Println("Context is Done")

	//========== Cancel context ==========

	// ctx := context.Background()
	// ctx , cancle := context.WithCancel(ctx)
	// go func() {
	// 	time.Sleep(1*time.Second) // waiting got task to done
	// 	fmt.Println(" task done | Canceling the context")
	// 	cancle()
	// }()
	
	// time.Sleep(3*time.Second)

	//========== context usage in logs ==========

	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", "12345")
	go logger(ctx, "Request started")
	time.Sleep(2*time.Second)
	go logger(ctx, "Request completed")
	time.Sleep(2*time.Second)
}