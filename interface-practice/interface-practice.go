package main

import "fmt"

// interfaces only contains methods - no attributes
type expense interface {
	cost() float64
}

type printer interface {
	print()
}

// collection of attributes that define an object i.e. human, dog
type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello world",
	}

	e.print()
	fmt.Println(e.cost())
}
