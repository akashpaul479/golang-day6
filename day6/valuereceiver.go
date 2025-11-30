package day6

import "fmt"

type student1 struct {
	Name  string
	Age   int
	Grade int
}

func (s student1) voting() {
	s.Age++
	fmt.Println("inside method:", s.Age)
}

func Valuepair() {
	s := student1{Name: "Akash", Age: 20}

	s.voting()
	fmt.Println("outside method:", s.Age)
}
