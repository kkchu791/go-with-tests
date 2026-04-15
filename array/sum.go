package main

// Slices instead
func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// []int{1, 2},
	// []int{0, 9}

	var sums []int
	println(sums)
	for _, nums := range numbersToSum {

		if len(nums) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
