package main

import (
	"context"
	"fmt"
)

func main() {
	const cap = 5
	ch := make(chan string, cap)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
		cancel()
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
			//time.Sleep(100 * time.Microsecond)
		}
	}

	close(ch)
	<-ctx.Done()
}
