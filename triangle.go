package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func CalcPerimeter(t Triangle) float64 {
	return t.Side * t.Side * t.Side
}

func CalcArea(t Triangle) float64 {
	return math.Sqrt(3) / 4 * t.Side * t.Side
}
