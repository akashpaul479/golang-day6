package day6

import "fmt"

type student struct {
	Name  string
	Age   int
	Marks int
	Grade string
}

func (s student) Greet() {
	fmt.Println("Hello!", s.Name)

}

func (s student) Isadult() bool {
	return s.Age >= 18

}
func Method() {
	s := student{Name: "Akash", Age: 20, Marks: 58, Grade: "A"}
	s.Greet()

	if s.Isadult() {
		fmt.Println(s.Name, "is an adult")
	} else {
		fmt.Println(s.Name, "is not an adult")
	}

}
