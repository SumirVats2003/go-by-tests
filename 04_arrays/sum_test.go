package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, sum int, expect int) {
		if sum != expect {
			t.Errorf("expected %q but got %q", expect, sum)
		}
	}

	t.Run("testing to sum an array of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		sum := Sum(nums)
		expect := 15
		checkSum(t, sum, expect)
	})

	t.Run("testing for collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}
		sum := Sum(nums)
		expect := 6
		checkSum(t, sum, expect)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("testing for intuitive case", func(t *testing.T) {
		sum := SumAll([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}
		checkSums(t, sum, expect)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("testing for intuitive case", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2}, []int{0, 9})
		expect := []int{2, 9}
		checkSums(t, sum, expect)
	})

	t.Run("testing for empty slice", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{9})
		expect := []int{0, 0}
		checkSums(t, sum, expect)
	})
}

func checkSums(t *testing.T, sum []int, expect []int) {
	t.Helper()
	if !slices.Equal(sum, expect) {
		t.Errorf("expected %q but got %q", expect, sum)
	}
}
