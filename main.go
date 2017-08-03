package main

import (
	"fmt"
)

func main() {
	controller := NewElevatorController(16, 10)
	controller.PickUp(0, 10)
	controller.PickUp(2, 5)
	for i := 0; i < 10; i++ {
		fmt.Printf("--------Step: %d--------", i)
		controller.Step()
		fmt.Printf("\n%v\n", controller.Status())
	}
}
