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

func SumAll(numsToSum ...[]int) []int {
	var length int = len(numsToSum)
	var result []int = make([]int, length)

	for i, nums := range numsToSum {
		result[i] = Sum(nums)
	}

	return result
}
