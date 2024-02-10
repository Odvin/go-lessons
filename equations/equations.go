package equations

import (
	"errors"
	"math"
)

// Solves the equation ax+b=0
func LinearEquation(a, b float64) (float64, error) {
	if a == 0 {
		return 0, errors.New("division by zero")
	}

	return -b / a, nil
}

// Solves the equation ax^2 + bx + c = 0
func QuadraticEquation(a, b, c float64) (float64, float64, error) {
	if a == 0 {
		return 0, 0, errors.New("division by zero")
	}

	d := b*b - 4*a*c
	if d < 0 {
		return 0, 0, errors.New("no real roots")
	}

	D := math.Sqrt(d)

	x1 := (-b + D) / (2 * a)
	x2 := (-b - D) / (2 * a)

	return x1, x2, nil
}
