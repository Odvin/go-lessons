package main

import (
	"fmt"
	"time"
)

func main() {
	arr := make([]int, 1e6)

	// The first and last elements of the slice
	fmt.Println(len(arr), arr[:5], arr[len(arr)-5:])

	// It will work correctly
	// Each goroutines works with one element
	// (two different goroutines do not operate with the same element)
	for i := 0; i < len(arr); i++ {
		go func(i int) {
			arr[i] = i * i
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(len(arr), arr[:5], arr[len(arr)-5:])
}
