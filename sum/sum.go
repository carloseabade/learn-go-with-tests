package main

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			numbers = append(numbers, 0)
		}
		sums = append(sums, Sum(numbers[1:]))
	}
	return sums
}
