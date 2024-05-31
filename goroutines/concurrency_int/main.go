package main

import "fmt"

// Problem with concurrency and not atomic operation
func main() {
	var count int

	for i := 0; i < 1e5; i++ {
		go func() {
			count++
		}()
	}

	// The result will not be equal 1e5
	fmt.Println(count)
}
