package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// wg.Add(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				fmt.Println("i:", i, "j:", j)
			}
		}(i)
	}

	wg.Wait()
}
