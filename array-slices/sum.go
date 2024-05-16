package main

func Sum(numbers []int) int {
	var sum int
	// _ is the blank identifier. Here, it omits the index
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// variadic functions take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {
	// use append to add the the return value of Sum
	// to the starting empty sums slice
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// create a function to add tails in all slices
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
