package main

import (
	"fmt"
	"sync"
)

func SplitChannel[T any](ch <-chan T, n int) []chan T {
	outs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outs[i] = make(chan T)
	}

	go func() {
		idx := 0
		for v := range ch {
			outs[idx] <- v
			idx = (idx + 1) % n
		}

		for _, ch := range outs {
			close(ch)
		}
	}()

	return outs
}

func main() {
	input := make(chan int)

	go func() {
		defer close(input)
		for i := 0; i < 100; i++ {
			input <- i
		}
	}()

	outs := SplitChannel(input, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range outs[0] {
			fmt.Println("ch0", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range outs[1] {
			fmt.Println("ch1", v)
		}
	}()

	wg.Wait()
}
