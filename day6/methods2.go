package day6

import "fmt"

type car1 struct {
	Name  string
	Price int
}

func (c *car1) Prices() {
	c.Price++
	fmt.Println("current price:", c.Price)
}

func Method3() {
	c := car1{Name: "Akash", Price: 250000}
	c.Prices()
	fmt.Println("updatedprice:", c.Price)
}
