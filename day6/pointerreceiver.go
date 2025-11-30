package day6

import "fmt"

type student2 struct {
	Name  string
	Age   int
	Grade int
}

func (s *student2) voting() {
	s.Age++
	fmt.Println("Inside method:", s.Age)
}

func Pointer() {
	s := student2{Name: "Akash", Age: 20}
	s.voting()
	fmt.Print("outside method:", s.Age)
}
