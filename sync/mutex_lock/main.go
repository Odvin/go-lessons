package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	var n int

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mutex.Lock()
			n++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(n)
}
