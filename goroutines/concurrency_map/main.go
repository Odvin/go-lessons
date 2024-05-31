package main

import "fmt"

// This example will not work in production but may work locally
// The problems with concurrent map write, read operations
func main() {
	m := make(map[int]int)

	go func() {
		for i := 0; i < 1000000000; i++ {
			// Concurrent write
			m[0] = i
		}
	}()

	go func() {
		for i := 0; i < 1000000000; i++ {
			//Concurrent read
			if m[0] < 0 {
				fmt.Println("Do something")
			}
		}
	}()
}
