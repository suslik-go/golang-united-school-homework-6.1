package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorNotEnoughCapacity = errors.New("There is no enough capacity")
	errorNotExistingIndex  = errors.New("There is no shape with such index")
	errorNotCircleFound    = errors.New("There is no circle in the shapes")
)

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
	//panic("implement me")
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("%w", errorNotEnoughCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//	panic("implement me")
	if i <= 0 || i > b.shapesCapacity {
		return nil, fmt.Errorf("%w", errorNotExistingIndex)
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	//	panic("implement me")
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("%w", errorNotExistingIndex)
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//	panic("implement me")
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("%w", errorNotExistingIndex)
	}

	b.shapes[i] = shape
	return shape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	//panic("implement me")
	var sum float64

	for _, val := range b.shapes {
		sum += val.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	//panic("implement me")
	var sum float64

	for _, val := range b.shapes {
		sum += val.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	//panic("implement me")
	var exists bool
	for i, val := range b.shapes {

		switch val.(type) {
		case *Circle:
			exists = true
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		}
	}

	if exists == false {
		return fmt.Errorf("%w", errorNotCircleFound)
	}
	return nil
}
