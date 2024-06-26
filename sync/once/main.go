package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("once processed")
			})
		}()
	}

	time.Sleep(time.Second)
}
