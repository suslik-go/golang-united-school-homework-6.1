package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func CalcPerimeter(c Circle) float64 {
	return math.Pi * c.Radius * 2
}

func CalcArea(c Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}
