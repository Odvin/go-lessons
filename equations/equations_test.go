package equations

import (
	"fmt"
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-3

func TestEquations(t *testing.T) {
	t.Run("LinearEquation", func(t *testing.T) {
		// ax + b=0
		a := 2.0
		b := -4.0

		x, _ := LinearEquation(a, b)

		assertAlmostRoot(t, a*x+b)
	})

	t.Run("LinearEquationError", func(t *testing.T) {
		// ax+b=0
		a := 0.0
		b := 3.0

		_, e := LinearEquation(a, b)

		if e != nil && e.Error() != "division by zero" {
			t.Errorf("expected 'division by zero' but got %g", e)
		}
	})

	t.Run("QuadraticEquation", func(t *testing.T) {
		// ax^2 + bx + c = 0
		// (x-2)(x-3) = 0 <=> x^2 - 5x + 6 = 0
		a := 1.0
		b := -5.0
		c := 6.0

		x1, x2, _ := QuadraticEquation(a, b, c)

		assertAlmostRoot(t, a*x1*x1+b*x1+c)
		assertAlmostRoot(t, a*x2*x2+b*x2+c)
	})

	t.Run("QuadraticEquationError", func(t *testing.T) {
		// ax^2 + bx + c = 0
		a := 0.0
		b := 3.0
		c := 1.0

		_, _, e := QuadraticEquation(a, b, c)

		if e != nil && e.Error() != "division by zero" {
			t.Errorf("expected 'division by zero' but got %g", e)
		}

		a = 1.0
		b = 1.0
		c = 5.0

		_, _, e = QuadraticEquation(a, b, c)

		if e != nil && e.Error() != "no real roots" {
			t.Errorf("expected 'no real roots' but got %g", e)
		}
	})
}

func assertAlmostRoot(t testing.TB, solution float64) {
	t.Helper()

	if math.Abs(solution) > float64EqualityThreshold {
		t.Errorf("expected 0 but got %.3f", solution)
	}
}

func ExampleLinearEquation() {
	// ax+b=0
	var a float64 = 2
	var b float64 = -4

	x, e := LinearEquation(a, b)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("x = %.1f", x)
	// Output: x = 2.0
}

func ExampleQuadraticEquation() {
	// ax^2 + bx + c = 0
	// (x-2)(x-3) = 0 <=> x^2 - 5x + 6 = 0
	var a float64 = 1.0
	var b float64 = -5.0
	var c float64 = 6.0

	x1, x2, e := QuadraticEquation(a, b, c)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("x1 = %.1f; x2 = %.1f", x1, x2)
	// Output: x1 = 3.0; x2 = 2.0
}
