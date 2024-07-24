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
	var result []int

	for _, nums := range numsToSum {
		result = append(result, Sum(nums))
	}

	return result
}

func SumAllTails(numsToSum ...[]int) []int {
	var result []int
	for _, nums := range numsToSum {
		var tail []int

		if len(nums) == 0 {
			tail = []int{0}
		} else {
			tail = nums[1:]
		}

		result = append(result, Sum(tail))
	}

	return result
}
