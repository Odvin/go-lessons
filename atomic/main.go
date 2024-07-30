package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int64

	for i := 0; i < 1000; i++ {
		go func() {
			// count++
			atomic.AddInt64(&count, 1)
		}()
	}

	var count2 atomic.Int64

	for i := 0; i < 1000; i++ {
		go func() {
			count2.Add(1)
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(count)
	fmt.Println(count2.Load())
}
