package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			// Mistake, we do not know what value if i will be printed
			fmt.Printf("i: '%d'\n", i)
		}()
	}

	for j := 0; j < 10; j++ {
		go func(j int) {
			// Will be printed all values of j but in the random order
			fmt.Printf("j: '%d'\n", j)
		}(j)
	}

	for k := 0; k < 10; k++ {
		k := k
		go func() {
			// Will be printed all values of k but in the random order
			fmt.Printf("k: '%d'\n", k)
		}()
	}

	for l := 0; l < 3; l++ {
		l := l
		go func() {
			for m := 0; m < 10; m++ {
				fmt.Printf("l: '%d', m: '%d'\n", l, m)
				// Randomize routine execution
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
}
