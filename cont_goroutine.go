package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func monitor2(ctx context.Context, number int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("monitor: %v, the received channel value is: %v, ending\n", number, v)
			return
		default:
			fmt.Printf("monitor: %v in progress...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}
func monitor1(ctx context.Context, number int) {
	for {
		go monitor2(ctx, number)
		select {
		case v := <-ctx.Done():
			// this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
			fmt.Printf("monitor: %v, the received channel value is: %v, ending\n", number, v)
			return
		default:
			fmt.Printf("monitor: %v in progress...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	var ctx context.Context = nil
	var cancel context.CancelFunc = nil
	ctx, cancel = context.WithCancel(context.Background())
	for i := 1; i <= 5; i = i + 1 {
		go monitor1(ctx, i)
	}
	time.Sleep(1 * time.Second)
	// close all gourtines
	cancel()
	// waiting 10 seconds, if the screen does not display <monitor: xxxx in progress>, all goroutines have been shut down
	time.Sleep(10 * time.Second)
	println(runtime.NumGoroutine())
	println("main program exit!!!!")
}
