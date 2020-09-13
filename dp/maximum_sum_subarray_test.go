package dp

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	want := 6
	if got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); got != want {
		t.Errorf("maxSubArray() = %d, want %d", got, want)
	}

	want = 1
	if got := maxSubArray([]int{1}); got != want {
		t.Errorf("maxSubArray() = %d, want %d", got, want)
	}

	want = 0
	if got := maxSubArray([]int{0}); got != want {
		t.Errorf("maxSubArray() = %d, want %d", got, want)
	}

	want = -1
	if got := maxSubArray([]int{-1}); got != want {
		t.Errorf("maxSubArray() = %d, want %d", got, want)
	}

	want = -2147483647
	if got := maxSubArray([]int{-2147483647}); got != want {
		t.Errorf("maxSubArray() = %d, want %d", got, want)
	}
}
