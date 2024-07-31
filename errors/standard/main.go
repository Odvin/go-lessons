package main

import (
	"errors"
	"fmt"
)

func intDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisor cannot be zero")
	}

	if b < 0 {
		return 0, fmt.Errorf("divisor %d has to be positive", b)
	}

	return a / b, nil
}

func main() {
	res, err := intDiv(4, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("intDiv result: %d\n", res)
}
