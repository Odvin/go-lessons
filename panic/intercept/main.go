package main

import (
	"fmt"
	"time"
)

func process() {

	go func() {
		defer func() {
			v := recover()
			fmt.Println("recovered:", v)
		}()
		panic("panic")
	}()

	time.Sleep(2 * time.Second)
}

func main() {
	process()

	fmt.Println("Done")
}
