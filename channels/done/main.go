package main

import (
	"fmt"
	"time"
)

func process(done chan struct{}) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			default:
				ch <- "processing"
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan struct{})

	// res := <-process(done)
	// fmt.Println(res)

	go func() {
		for v := range process(done) {
			fmt.Println(v)
		}
	}()

	time.Sleep(3 * time.Second)
	close(done)
}
