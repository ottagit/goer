package main

var sum int

func Sum(numbers [6]int) int {
	for _, number := range numbers {
		sum += number
	}
	return sum
}
