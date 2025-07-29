package main

import (
	"testing"
	"reflect"
)

/*
(From Go 1.21, slices standard package is available, which has slices.
Equal function to do a simple shallow compare on slices, 
where you don't need to worry about the types like the above case. 
Note that this function expects the elements to be comparable. 
So, it can't be applied to slices with non-comparable elements like 2D slices.)
*/

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Sum all, takes a varying number of slices, returning a new slice containing the totals for each slice passed in. 2*2", func(t *testing.T) {
		numbers := []int{1,2}
		numbers2 := []int{0,9}

		got := SumAllFixed(numbers, numbers2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Sum all, takes a varying number of slices, returning a new slice containing the totals for each slice passed in. 2*2", func(t *testing.T) {
		numbers := []int{1, 1, 1}

		got := SumAllDynamic(numbers)
		want := []int{3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("make the sums of some slices, dynamic slice", func(t *testing.T) {
		got := SumAllTailsDynamic([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("make the sums of some slices, fixed slice", func(t *testing.T) {
		got := SumAllTailsFixed([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices, dynamic slice", func(t *testing.T) {
		got := SumAllTailsDynamic([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices, fixed slice", func (t *testing.T) {
		got := SumAllTailsFixed([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}) 
}