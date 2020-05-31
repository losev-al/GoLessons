package main

import "fmt"

func demo() {
	a := 1
	b := 1
	c := 1
	println(&a)
	fmt.Println(&b)
	myPrint(&c)
}

func myPrint(p *int) {
	println(p)
}

func main() {
	demo()
	d := 1
	println(&d)
}
