// One of the other uses of context is to terminate
// all goroutines spawned by a function once the context is cancelled.
// This is done by passing/propagating the context to all child goroutines.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type output struct {
	count int
	err   error
}

func dbTaskA(ctx context.Context, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("dbTaskA was terminated with error:", ctx.Err())
		return 0, fmt.Errorf("dbTaskA error: %w", ctx.Err())
	case <-time.After(7 * time.Second):
		return 10, nil
	}
}

func dbTaskB(ctx context.Context, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("dbTaskB was terminated with error:", ctx.Err())
		return 0, fmt.Errorf("dbTaskB error %w", ctx.Err())
	case <-time.After(5 * time.Second):
		return 20, nil
	}
}

func webApi(ctx context.Context) (int, error) {
	wg := sync.WaitGroup{}

	outputChanA := make(chan output)
	outputChamB := make(chan output)

	wg.Add(1)
	go func() {
		count, err := dbTaskA(ctx, &wg)
		outputChanA <- output{
			count: count,
			err:   err,
		}
		close(outputChanA)
	}()

	wg.Add(1)
	go func() {
		count, err := dbTaskB(ctx, &wg)
		outputChamB <- output{
			count: count,
			err:   err,
		}
		close(outputChamB)
	}()

	wg.Wait()

	outputA := <-outputChanA
	if outputA.err != nil {
		return 0, outputA.err
	}

	outputB := <-outputChamB
	if outputB.err != nil {
		return 0, outputB.err
	}

	output := outputA.count + outputB.count

	return output, nil
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	count, err := webApi(ctx)
	if err != nil {
		fmt.Println("webApi call exited with error:", err)
		return
	}

	fmt.Println("Count is:", count)
}
