package main

import (
	"fmt"
	"strconv"
)

type FizzBuzzMap struct {
	div_int         int
	replacement_str string
}

func FizzBuzz(n int) []string {
	result := []string{}
	mapVals := []FizzBuzzMap{
		{3, "Fizz"},
		{5, "Buzz"},
	}

	for i := 1; i <= n; i++ {
		currStr := ""
		for _, mapVal := range mapVals {
			if i%mapVal.div_int == 0 {
				currStr += mapVal.replacement_str
			}
		}

		if currStr == "" {
			currStr += strconv.Itoa(i)
		}
		result = append(result, currStr)
	}

	return result
}

func main() {
	fmt.Printf("%v", FizzBuzz(3))
}
