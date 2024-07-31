package main

import (
	"errors"
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func (e *areaError) radiusNegative() bool {
	return e.radius < 0
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			err:    "radius is negative",
			radius: radius,
		}
	}

	return math.Pi * radius * radius, nil
}

func main() {
	area, err := circleArea(-5)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			if areaError.radiusNegative() {
				fmt.Printf("Area calculation failed, radius %0.2f is less than zero\n", areaError.radius)
			}
			fmt.Println(err)
			return
		}
		return
	}

	fmt.Printf("Area of rectangle %0.2f\n", area)
}
