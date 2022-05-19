package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func monitor(ctx context.Context, index int) {
	for {
		select {
		case <-ctx.Done():
			// this branch is only reached when the ch channel is closed, or when data is sent(either true or false)
			fmt.Printf("monitor %v, end of monitoring. \n", index)
			return
		default:
			var value interface{} = ctx.Value("Nets")
			fmt.Printf("monitor %v, is monitoring %v\n", index, value)
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	var ctx01 context.Context = nil
	var ctx02 context.Context = nil
	var cancel context.CancelFunc = nil
	ctx01, cancel = context.WithCancel(context.Background())
	ctx02, cancel = context.WithTimeout(ctx01, 1*time.Second)
	var ctx03 context.Context = context.WithValue(ctx02, "Nets", "Champion") // key: "Nets", value: "Champion"

	defer cancel()
	for i := 1; i <= 5; i = i + 1 {
		go monitor(ctx03, i)
	}
	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("the cause of cancel is: ", ctx02.Err())
	}
	println(runtime.NumGoroutine())
	println("main program exit!!!!")
}
