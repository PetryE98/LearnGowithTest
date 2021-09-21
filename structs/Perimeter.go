package structs

import "math"

func Perimeter(Rectangle rectangle) float64 {

	return 2 * (Rectangle.width + Rectangle.height)

}

func Area(Rectangle rectangle) float64 {

	return Rectangle.height * Rectangle.width
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {

	return 0.5 * t.Base * t.Height
}
