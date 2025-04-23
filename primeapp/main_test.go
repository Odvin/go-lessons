package main

import "testing"

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
