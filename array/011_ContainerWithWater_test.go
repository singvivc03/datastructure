package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	want := 49
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	if got := maxArea(height); got != want {
		t.Errorf("maxArea(%v) = %d, want %d", height, got, want)
	}
}
