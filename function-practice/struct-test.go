package main

import "fmt"

type Car struct {
	Make  string
	Model string
}

func printCar(car Car) {
	fmt.Printf("This car is a %v %v", car.Make, car.Model)
}
