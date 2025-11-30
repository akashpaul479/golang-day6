package day6

import "fmt"

type car2 struct {
	Name  string
	Price int
}

func (c *car2) updatedprice(newprice int) {
	c.Price = newprice
}
func Method2() {
	c := car2{Name: "Thar", Price: 25000}
	c.updatedprice(30000)
	fmt.Println(c.Price)
}
