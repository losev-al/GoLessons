package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan<- struct{}) {
		fmt.Println("signal func start")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("sleep ended")
		close(ch)
	}(channel)

	reader := func(index int, ch <-chan struct{}) {
		for true {
			select {
			case <-ch:
				fmt.Println(index, " closed")
				wg.Done()
				return
			default:
				fmt.Println(index, "do work")
				time.Sleep(30 * time.Millisecond)
			}
		}
	}

	go reader(1, channel)
	go reader(2, channel)

	fmt.Println("prog finish")
	wg.Wait()
}
