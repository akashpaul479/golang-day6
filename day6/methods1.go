package day6

import "fmt"

type car struct {
	Name    string
	Modelno int
	Partno  int
}

func (c car) Greet() {
	fmt.Println("hey! iam a brand , name", c.Name)
}

func Method1() {
	c := car{Name: "Thar", Modelno: 56354, Partno: 43625}
	c.Greet()
	fmt.Println(c.Name, c.Modelno, c.Partno)
}
