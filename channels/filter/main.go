package main

import "fmt"

func Filter[T any](in <-chan T, condition func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			if condition(v) {
				out <- v
			}
		}
	}()

	return out
}

func main() {
	in := make(chan int)

	go func() {
		defer close(in)
		for i := 0; i < 15; i++ {
			in <- i
		}
	}()

	isOdd := func(n int) bool {
		return n%2 != 0

	}

	for v := range Filter(in, isOdd) {
		fmt.Println(v)
	}
}
