package main

import "fmt"

func multiple(a, b int) (result int) {
	return a * b
}

func bar() {
	fmt.Println(multiple(2, 3))
}
