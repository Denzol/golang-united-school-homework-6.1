package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	Perimeter := 2 * math.Pi * c.Radius
	return Perimeter
}

func (c Circle) CalcArea() float64 {
	Area := math.Pi * c.Radius * c.Radius
	return Area
}
