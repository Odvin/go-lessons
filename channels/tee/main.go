package main

import (
	"fmt"
	"sync"
)

func Tee[T any](input <-chan T, n int) []chan T {
	outs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outs[i] = make(chan T)
	}

	go func() {
		for v := range input {
			for i := 0; i < n; i++ {
				outs[i] <- v
			}
		}

		for i := 0; i < n; i++ {
			close(outs[i])
		}
	}()

	return outs
}

func main() {
	input := make(chan int)

	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()

	outs := Tee(input, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for v := range outs[0] {
			fmt.Println("Tee0", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range outs[1] {
			fmt.Println("Tee1", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range outs[2] {
			fmt.Println("Tee2", v)
		}
	}()

	wg.Wait()
}
