package arrayandslices

import (
	"slices"
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		// On pourrait tres bien utiliser la function SumSlice dans ce cas.
		got := SumSlices(numbers)

		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlices(numbers)

		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Sum all with make", func(t *testing.T) {
		got := SumAllWithMake([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum all with append", func(t *testing.T) {
		got := SumAllWithAppend([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
