package array

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	want := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	if got := fourSum(nums, target); !isMatching(got, want) {
		t.Errorf("fourSum(%v, %d)=%v, want %v", nums, target, got, want)
	}
}

func isMatching(actual [][]int, expected [][]int) bool {
	expectedLength, actualLength := len(expected), len(actual)
	if expectedLength != actualLength {
		return false
	}
	for i := 0; i < expectedLength; i++ {
		currentExpected := expected[i]
		currentActual := actual[i]
		if (currentExpected[0] != currentActual[0]) || (currentExpected[1] != currentActual[1]) ||
			(currentExpected[2] != currentActual[2]) || (currentExpected[3] != currentActual[3]) {
			return false
		}
	}
	return true
}
