package miniproject

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
func readFloat(prompt string) float64 {
	for {
		Text := ReadLine(prompt)
		if Text == "" {
			fmt.Println("Please enter a number")
			continue
		}
		v, err := strconv.ParseFloat(Text, 64)
		if err != nil {
			fmt.Println("Invalid number")
			continue
		}
		return v
	}
}

var reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func Shapecalculator() {
	for {
		fmt.Println("__Shape calculator__")
		fmt.Println("1.rectangle:")
		fmt.Println("2.circle:")
		fmt.Println("3.square:")
		fmt.Println("4.Exit:")

		choice := ReadLine("Entyer your choice:")
		var shape1 shape1
		switch choice {
		case "1":
			w := readFloat("Enter width:")
			h := readFloat("Enter height:")
			shape1 = rectangle1{Width: w, Height: h}

		case "2":
			r := readFloat("Enter radius:")
			shape1 = Circle1{Radius: r}
		case "3":
			s := readFloat("Enter side:")
			shape1 = square1{Side: s}
		case "4":
			fmt.Println("Good bye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
		fmt.Println("\n_____________________________________")
		fmt.Printf("Area :%.2f | Perimeter :%.2f\n", shape1.Area(), shape1.Perimeter())

	}
}
