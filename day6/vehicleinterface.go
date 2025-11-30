package day6

import "fmt"

type vehicle interface {
	start()
	stop()
}

type bike struct{}
type cars struct{}

func (b bike) start() {
	fmt.Println("Bike started!")

}
func (b bike) stop() {
	fmt.Println("Bike stopped!")
}

func (c cars) start() {
	fmt.Println("car started!")
}
func (c cars) stop() {
	fmt.Println("Car stopped!")
}
func Vehicleinterface() {
	var v1 vehicle = bike{}
	var v2 vehicle = cars{}

	v1.start()
	v1.stop()
	v2.start()
	v2.stop()
}
