package main

func Sum(arr []int) int {
	var result int = 0
	// for i := 0; i < 5; i++ {
	// 	result += arr[i]
	// }

	for _, num := range arr {
		result += num
	}

	return result
}
