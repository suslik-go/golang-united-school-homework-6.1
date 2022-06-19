package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func CalcPerimeter(r Rectangle) float64 {
	return 2*r.Height + 2*r.Weight
}

func CalcArea(r Rectangle) float64 {
	return r.Height * r.Weight
}
