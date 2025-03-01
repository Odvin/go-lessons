package main

import "fmt"

func Transformer[T any](in <-chan T, action func(T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			out <- action(v)
		}
	}()

	return out
}

func main() {
	in := make(chan int)

	go func() {
		defer close(in)
		for i := 1; i < 10; i++ {
			in <- i
		}
	}()

	action := func(i int) int {
		return i * i
	}

	for v := range Transformer(in, action) {
		fmt.Println(v)
	}
}
