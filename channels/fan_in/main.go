package main

import (
	"fmt"
	"sync"
)

func MergeChannels[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	output := make(chan T)
	for _, ch := range channels {
		go func(ch <-chan T) {
			defer wg.Done()
			for value := range ch {
				output <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	channel01 := make(chan int)
	channel02 := make(chan int)
	channel03 := make(chan int)

	go func() {
		defer func() {
			close(channel01)
			close(channel02)
			close(channel03)
		}()

		for i := 0; i < 100; i += 3 {
			channel01 <- i
			channel02 <- i + 1
			channel03 <- i + 2
		}
	}()

	for value := range MergeChannels(channel01, channel02, channel03) {
		fmt.Println(value)
	}
}
