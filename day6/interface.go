package day6

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r rectangle) area() float64 {
	return r.Height * r.Width
}
func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func Printarea(s shape) {
	fmt.Println("Shape:", s.area())
}
func Interface() {
	r := rectangle{Width: 3, Height: 5}
	c := Circle{Radius: 5}
	Printarea(r)
	Printarea(c)
}
