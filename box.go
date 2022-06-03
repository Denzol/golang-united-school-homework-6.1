package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return errors.New("No place for it!")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Value is incorrect!")
	}
	Shape := b.shapes[i]
	if Shape != nil {
		return Shape, nil
	} else {
		return Shape, errors.New("Value is empty!")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes) {
		return nil, errors.New("Value is incorrect!")
	}
	Shape := b.shapes[i]
	if Shape != nil {
		b.shapes[i] = b.shapes[len(b.shapes)-1]
		b.shapes[len(b.shapes)-1] = nil
		b.shapes = b.shapes[0 : len(b.shapes)-1]
		return Shape, nil
	} else {
		return Shape, errors.New("Value is empty!")
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes) {
		return nil, errors.New("Value is incorrect!")
	}
	Shape := b.shapes[i]
	if Shape != nil {
		b.shapes[i] = shape
		return shape, nil
	} else {
		return Shape, errors.New("Value is empty!")
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for i := 0; i < len(b.shapes); i++ {
		shape := b.shapes[i]
		sumPerimeter = sumPerimeter + shape.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for i := 0; i < len(b.shapes); i++ {
		shape := b.shapes[i]
		sumArea = sumArea + shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var count int
	for i := 0; i < len(b.shapes); i++ {
		shape := b.shapes[i]
		_, ok := shape.(Circle)
		if ok {
			b.shapes[i] = b.shapes[len(b.shapes)-1]
			b.shapes[len(b.shapes)-1] = nil
			b.shapes = b.shapes[0 : len(b.shapes)-1]
			count++
		}
	}
	if count == 0 {
		return errors.New("No Circles!")
	}
	return nil
}
