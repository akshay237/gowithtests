package goarr_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

func TestSumSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	got := SumSlice(nums)
	want := 21
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}

// Slice can only compared with nil
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	checkSumAssert(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	// if slices are not empty
	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		checkSumAssert(t, got, want)
	})

	// some slices can be empty
	t.Run("make sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 2, 7})
		want := []int{0, 9}
		checkSumAssert(t, got, want)
	})
}

func checkSumAssert(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}

// to make an slice from an array use [:] eg: sli := arr[:]
// sli points to same location of the array where arr points so if any update in sli will update arr also if we have used copy then it won't
