package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 6 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			// Sometimes, it is useful to print the inputs to the function in the error message. Here, the
			// %v placeholder is used to print the default format, which works well for arrays
			t.Errorf("got %d want %d, given, %v", got, want, numbers)
		}
	})
}
