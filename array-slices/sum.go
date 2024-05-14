package main

var sum int

func Sum(numbers [6]int) int {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
