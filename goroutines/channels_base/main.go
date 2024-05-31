package main

import "fmt"

func generate(in chan<- int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
	// If the chanel will not be closed it will trigger deadlock error
	close(in)
}

func square(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}

	close(out)
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go generate(in)
	go square(in, out)

	for i := range out {
		fmt.Println(i)
	}

}
