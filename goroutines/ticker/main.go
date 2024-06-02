package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)

	// Reading for nte ticker channel
	for t := range ticker.C {
		fmt.Println(t)
	}
}
