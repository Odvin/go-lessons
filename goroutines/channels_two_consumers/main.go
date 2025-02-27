package main

import (
	"fmt"
	"sync"
)

func producer(c chan int) {
	for i := 0; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func consumer(id int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range c {
		fmt.Printf("Cunsumer: %d, value: %d\n", id, i)
	}
}

func main() {
	var c = make(chan int)
	var wg sync.WaitGroup

	go producer(c)

	wg.Add(2)
	go consumer(1, c, &wg)
	go consumer(2, c, &wg)
	wg.Wait()

	fmt.Println("The end")
}
