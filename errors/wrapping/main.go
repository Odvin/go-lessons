package main

import (
	"errors"
	"fmt"
)

var noRows = errors.New("no rows found")

func getRecord() error {
	return noRows
}

func webService() error {
	if err := getRecord(); err != nil {
		return fmt.Errorf("%w when calling db", err)
	}

	return nil
}

func main() {
	if err := webService(); err != nil {
		if errors.Is(err, noRows) {
			fmt.Printf("The searched record was not found. DB error: %s\n", err)
		}
		fmt.Println("unknown error when searching the record")
	}

	fmt.Println("webservice call successful")
}
