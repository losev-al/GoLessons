package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	go func(ch chan bool) {
		select {
		case <-printAndSleep(fmt.Sprint(time.Now(), " 2 seconds"), 2*time.Second):
			println(fmt.Sprint(time.Now(), " 2 seconds finish"))
		case <-printAndSleep(fmt.Sprint(time.Now(), " 1 second"), 1*time.Second):
			println(fmt.Sprint(time.Now(), " 1 second finish"))
		}
		ch <- true
	}(channel)
	<-channel
}

func printAndSleep(message string, sleepTime time.Duration) chan struct{} {
	println(message)
	channel := make(chan struct{})
	go func(ch chan struct{}) {
		time.Sleep(sleepTime)
		close(channel)
	}(channel)
	return channel
}
