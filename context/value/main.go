package main

import (
	"context"
	"fmt"
)

func stepA(ctx context.Context) {
	fmt.Println("StepA. Context value for keyA:", ctx.Value("keyA"))
	anotherCtx := context.WithValue(ctx, "keyB", "valB")
	stepB(anotherCtx)

}

func stepB(ctx context.Context) {
	fmt.Println("StepB. Context value from keyA:", ctx.Value("keyA"))
	fmt.Println("StepB. Context value from keyB:", ctx.Value("keyB"))
}

func main() {
	ctx := context.WithValue(context.Background(), "keyA", "valA")
	stepA(ctx)
}
