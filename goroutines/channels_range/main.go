package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}

		// If the channel is not closed receive deadlock
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
