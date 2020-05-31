package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	emps := 20
	//ch := make(chan string, emps)
	ch := make(chan string)

	for e := 0; e < emps; e++ {
		go func(index int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- fmt.Sprint("paper ", index)
			fmt.Println(index)
		}(e)
	}

	for emps > 0 {
		time.Sleep(50 * time.Millisecond)
		p := <-ch
		fmt.Println(p)
		emps--
	}
}
