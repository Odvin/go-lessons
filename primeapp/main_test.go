package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		message  string
	}{
		{"prime", 7, true, "7 is prime number"},
		{"not prime", 8, false, "8 is not prime because it is divisible by 2"},
		{"not prime", 0, false, "0 is not prime, by definition"},
		{"not prime", -3, false, "negative numbers are not prime, by definition"},
	}

	for _, tt := range primeTests {
		result, msg := isPrime(tt.testNum)

		if result != tt.expected {
			t.Errorf("isPrime(%d) :: result expected: %v, actual: %v", tt.testNum, tt.expected, result)
		}

		if msg != tt.message {
			t.Errorf("isPrime(%d) :: message expected: %s, actual: %s", tt.testNum, tt.message, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if string(out) != "--> " {
		t.Errorf("prompt() output expected: --> , actual: %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Enter a whole number to check if it is prime. Enter q to quit.") {
		t.Errorf("intro text is not correct; got: %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	checkNumbersTests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a whole number"},
		{"prime", "7", "7 is prime number"},
		{"negative", "-3", "negative numbers are not prime, by definition"},
		{"not prime", "8", "8 is not prime because it is divisible by 2"},
		{"exit", "q", ""},
	}

	for _, tt := range checkNumbersTests {
		input := strings.NewReader(tt.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, tt.expected) {
			t.Errorf("checkNumbers() failed on %s test; expected: %s, got: %s", tt.name, tt.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
