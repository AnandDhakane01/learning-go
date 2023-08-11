package learninggo

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area(circle Circle) float64 {
	return math.Pi * circle.Radius * circle.Radius
}

type Shape interface {
	Area() float64
}