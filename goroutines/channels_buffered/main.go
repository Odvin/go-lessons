package main

import "fmt"

func generate(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	c := make(chan int)

	// Without goroutines receive deadlock
	// generate(c)

	go generate(c)

	for x := range c {
		fmt.Println(x)
	}

	cb := make(chan int, 5)

	generate(cb)

	for y := range cb {
		fmt.Println(y)
	}
}
