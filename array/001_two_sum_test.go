package array

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	nums := []int{2, 7, 11, 15}
	target := 9

	if got := twoSum(nums, target); got[0] != want[0] && got[1] != want[1] {
		t.Errorf("twoSum(%v, %d) = %v, want %v", nums, target, got, want)
	}

	want = []int{1, 2}
	nums = []int{3, 2, 4}
	target = 6

	if got := twoSum(nums, target); got[0] != want[0] && got[1] != want[1] {
		t.Errorf("twoSum(%v, %d) = %v, want %v", nums, target, got, want)
	}

	want = []int{0, 1}
	nums = []int{3, 3}
	target = 6

	if got := twoSum(nums, target); got[0] != want[0] && got[1] != want[1] {
		t.Errorf("twoSum(%v, %d) = %v, want %v", nums, target, got, want)
	}
}
