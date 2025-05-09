package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan
	close(doneChan)

	fmt.Println("Goodbye!")
}

func readUserInput(r io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(r)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number to check if it is prime. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("--> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("negative numbers are not prime, by definition")
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime number", n)
}
