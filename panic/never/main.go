package main

import (
	"fmt"
	"log"
	"time"
)

func task() {
	time.Sleep(time.Second)
	panic("uninfected situation")
}

func NeverExit(name string, action func()) {
	defer func() {
		if v := recover(); v != nil {
			log.Println(name, "is crashed - restarting...")
			go NeverExit(name, action)
		}
	}()

	if action != nil {
		action()
	}
}

func main() {
	go NeverExit("first task", task)
	go NeverExit("second task", task)

	time.Sleep(6 * time.Second)

	fmt.Println("Done")
}
