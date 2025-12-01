package day6

import (
	"fmt"
	"math"
)

type shape1 interface {
	Area() float64
	Perimeter() float64
}
type rectangle1 struct {
	Width  float64
	Height float64
}
type Circle1 struct {
	Radius float64
}
type square1 struct {
	Side float64
}

func (r rectangle1) Area() float64 {
	return r.Width * r.Height
}
func (r rectangle1) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle1) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle1) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func (s square1) Area() float64 {
	return s.Side * s.Side
}
func (s square1) Perimeter() float64 {
	return 4 * s.Side
}
func Printshape(s shape1) {
	fmt.Printf("Area :%.2f | Perimeter :%.2f\n", s.Area(), s.Perimeter())
}
func Shapeareacalculator() {
	r := rectangle1{Width: 3, Height: 4}
	c := Circle1{Radius: 5}
	s := square1{Side: 4}

	shapes := []shape1{r, c, s}

	for _, shape := range shapes {
		Printshape(shape)
	}

}
