package main

import (
	"fmt"
	"time"
)

type Printer struct{}

func printHello() {
	fmt.Println("From named function")
}

func (Printer) printHello() {
	fmt.Println("From struct method")
}

func main() {
	go func() {
		fmt.Println("From anonymous function")
		time.Sleep(time.Second * 2)
		fmt.Println("Will not be printed")
	}()

	go printHello()

	var p Printer
	go p.printHello()

	time.Sleep(time.Second)
}
