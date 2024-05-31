package main

import "fmt"

func main() {
	first := make(chan int)
	prev := first

	const goroutinesCount = 10000

	for i := 0; i < goroutinesCount; i++ {
		next := make(chan int)

		go func(prev chan int) {
			number := <-prev
			next <- number

			//next <- <-prev
		}(prev)

		prev = next
	}

	// 42 <- [c] <- ... [c] <- [c] <- 42
	first <- 42

	fmt.Println(<-prev)
}
