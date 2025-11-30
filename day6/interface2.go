package day6

import "fmt"

type animal interface {
	speak() string
}

type dog struct{}
type cat struct{}
type cow struct{}

func (d dog) speak() string {
	return "bark!"
}
func (c cat) speak() string {
	return "meow!"
}
func (co cow) speak() string {
	return "moo!"
}
func Animalsinterface() {
	var a1 animal = dog{}
	var a2 animal = cat{}
	var a3 animal = cow{}

	fmt.Println("Dog:", a1.speak())
	fmt.Println("Cat:", a2.speak())
	fmt.Print("Cow:", a3.speak())
}
