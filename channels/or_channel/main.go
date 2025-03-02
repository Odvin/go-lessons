package main

import (
	"fmt"
	"time"
)

func or[T any](input ...<-chan T) <-chan T {
	switch len(input) {
	case 0:
		return nil
	case 1:
		return input[0]
	}

	out := make(chan T)

	go func() {
		defer close(out)
		switch len(input) {
		case 2:
			select {
			case <-input[0]:
			case <-input[1]:
			}
		default:
			select {
			case <-input[0]:
			case <-input[1]:
			case <-input[2]:
			case <-or(input[3:]...):
			}
		}
	}()

	return out
}

func main() {
	start := time.Now()

	<-or(
		time.After(2*time.Second),
		time.After(5*time.Second),
	)

	fmt.Printf("called after: %s", time.Since(start))
}
