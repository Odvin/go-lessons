package main

import (
	"fmt"
	"time"
)

type Worker struct {
	done    chan struct{}
	stopped chan struct{}
}

func NewWorker(message string) Worker {
	worker := Worker{
		done:    make(chan struct{}),
		stopped: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(time.Second)

		defer func() {
			ticker.Stop()
			close(worker.stopped)
		}()

		for {
			select {
			case <-worker.done:
				return
			case <-ticker.C:
				fmt.Println(message)
			}
		}
	}()

	return worker
}

func (w *Worker) Shutdown() {
	close(w.done)
	fmt.Println("worker is stopping")
	<-w.stopped
	fmt.Println("worker is stopped")
}

func main() {
	worker := NewWorker("working")
	time.Sleep(3 * time.Second)
	worker.Shutdown()
}
