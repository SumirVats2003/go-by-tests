package arrays

func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func SumAll(lists ...[]int) []int {
	var sums []int

	for _, numbers := range lists {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(lists ...[]int) []int {
	var sums []int

	for _, numbers := range lists {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
