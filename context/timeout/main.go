package main

import (
	"context"
	"fmt"
	"time"
)

func longRunning(ctx context.Context) (int, error) {
	count := 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			count += i
			fmt.Println("Current value of the count:", count)
			time.Sleep(2 * time.Second)
		}
	}
	return count, nil
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	count, err := longRunning(ctx)
	if err != nil {
		fmt.Println("Long running task exits with error:", err)
	}

	fmt.Println("Coint is:", count)
}
