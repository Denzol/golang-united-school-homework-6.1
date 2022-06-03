package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	Perimeter := 3 * t.Side
	return Perimeter
}

func (t Triangle) CalcArea() float64 {
	Area := t.Side * t.Side * math.Sqrt(3) * 0.25
	return Area
}
