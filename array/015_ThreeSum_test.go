package array

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	want := [][]int{{-1, 0, 1}, {-1, -1, 2}}
	nums := []int{-1, 0, 1, 2, -1, -4}
	if got := threeSum(nums); !isMatch(got, want) {
		t.Errorf("threeSum(%v)=%v, want %v", nums, got, want)
	}

	want = [][]int{}
	nums = []int{}
	if got := threeSum(nums); !isMatch(got, want) {
		t.Errorf("threeSum(%v)=%v, want %v", nums, got, want)
	}

	want = [][]int{}
	nums = []int{0}
	if got := threeSum(nums); !isMatch(got, want) {
		t.Errorf("threeSum(%v)=%v, want %v", nums, got, want)
	}
}

func isMatch(got [][]int, want [][]int) bool {
	for i := 0; i < len(want); i++ {
		actualRow, expectedRow := got[i], want[i]
		if actualRow[0] != expectedRow[0] || actualRow[1] != expectedRow[1] || actualRow[2] != expectedRow[2] {
			return false
		}
	}
	return true
}
