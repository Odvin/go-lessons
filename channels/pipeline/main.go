package main

import "fmt"

func generate[T any](values ...T) chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()

	return ch
}

func process[T any](in <-chan T, action func(T) T) chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for v := range in {
			ch <- action(v)
		}
	}()

	return ch
}

func main() {
	vales := []int{1, 2, 3, 4, 5}

	mul := func(i int) int {
		return i * i
	}

	for v := range process(generate(vales...), mul) {
		fmt.Println(v)
	}
}
