package main

import "fmt"

func concat(str1, str2 string) string {
	return str1 + str2
}

func main() {
	myCar := Car{"Honda", "CRV"}

	fmt.Println(concat("hello", " world"))
	bar()
	printCar(myCar)
}
