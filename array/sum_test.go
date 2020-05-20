package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		wanted := 6
		if got != wanted {
			t.Errorf("wrong result")
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	wanted := []int{3, 9}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("Wanted %v got %v", wanted, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum sum numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		wanted := []int{2, 9}
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Wanted %v got %v", wanted, got)
		}
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		wanted := []int{0, 9}
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Wanted %v got %v", wanted, got)
		}
	})

}
