package main

import (
	"testing"
	"reflect"
)

func TestArray(t *testing.T) {

	t.Run("Should pass with 5 size of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		actual := Sum(numbers) // t.Fatal("not implemented")

		if expected != actual {
			// not %i, it is %d
			t.Errorf("TOBE : '%d', but ASIS : '%d', given '%v'", expected, actual, numbers)
		}
	})
	t.Run("Should pass with 4 size of integer array", func(t *testing.T) {
		numbers := []int{1,2,3,4}
		expected := 10
		actual := Sum(numbers)
		if expected != actual {
			// not %i, it is %d
			t.Errorf("TOBE : '%d', but ASIS : '%d', given '%v'", expected, actual, numbers)
		}
	})

	t.Run("Should return a new slice containing the totals for each slice pass in", func(t *testing.T) {
		slice1 := []int{1,2,3}
		slice2 := []int{0,9}
		expected := []int{6,9}
		actual := SumAll(slice1, slice2)

		// Go does not let you use equality operators with slice.
		// We can use reflect.DeepEqual() which is useful
		// for seeing if ANY two variables are the same.
		// NOTE: reflect.DeepEqual() is NOT type safe!
		// you can even compare slices and string!
		// so please be careful when using it.
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("TOBE : '%v', but ASIS : '%v'", expected, actual)
		}
	})

	t.Run("Should return a new slice containing the totals for 1 slice pass in", func(t *testing.T) {
		expected := []int{10}
		actual := SumAll([]int{1,2,3,4})

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("TOBE : '%v', but ASIS : '%v'", expected, actual)
		}
	})
}

func TestSumAllTail(t *testing.T){

	t.Run("Should return a new slice containing tail elements for each slice pass in", func(t *testing.T) {
		expected := []int{2,9}
		actual := SumAllTail([]int{1,2}, []int{0,1,9})
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected : '%v', but Actual is : '%v'", expected, actual)
		}
	})

	t.Run("Should return a new slice containing tail element for 1 slice pass in", func(t *testing.T) {
		expected := []int{2}
		actual := SumAllTail([]int{1,2})
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected : '%v', but Actual is : '%v'", expected, actual)
		}
	})

	t.Run("Should return a new slice containing tail element for 1 slice and zero slice pass in", func(t *testing.T) {
		expected := []int{2}
		actual := SumAllTail([]int{1,2}, []int{})
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected : '%v', but Actual is : '%v'", expected, actual)
		}
	})
}
