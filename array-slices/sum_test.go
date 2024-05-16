package main

import (
	"reflect"
	"testing"
)

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

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		got := SumAllTails(slice1, slice2)
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 9}

		got := SumAllTails(slice1, slice2)
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
