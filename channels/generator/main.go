package main

import "fmt"

func Generator(start, end int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; i <= end; i++ {
			out <- i
		}
	}()

	return out
}

func main() {
	for v := range Generator(1, 5) {
		fmt.Println(v)
	}
}
