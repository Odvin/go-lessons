package main

import (
	"fmt"
	"sync"
)

func parse(in <-chan string) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for v := range in {
			ch <- v + "+++"
		}
	}()

	return ch
}

func send(in <-chan string, n int) chan string {
	ch := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for v := range in {
				ch <- fmt.Sprintf("worker %d: %s", w, v)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	input := make(chan string)

	go func() {
		defer close(input)
		for i := 0; i < 20; i++ {
			input <- "value"
		}
	}()

	for v := range send(parse(input), 10) {
		fmt.Println(v)
	}
}
