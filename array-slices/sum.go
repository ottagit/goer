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
	lengthOfNumbers := len(numbersToSum)
	// create a slice with a starting capacity equal to the length of
	// the numbersToSum
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
