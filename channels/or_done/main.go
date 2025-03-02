package main

import (
	"fmt"
	"time"
)

func orDone[T any](input <-chan T, done chan struct{}) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case value, open := <-input:
				if !open {
					return
				}
				out <- value
			}
		}
	}()

	return out
}

func main() {
	ch := make(chan string)

	go func() {
		for {
			ch <- "working"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	for v := range orDone(ch, done) {
		fmt.Println(v)
	}
}
