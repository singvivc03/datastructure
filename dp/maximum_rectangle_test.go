package dp

import (
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	want := 6
	matrix := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	if got := maximalRectangle(matrix); got != want {
		t.Errorf("maximalRectangle(%v) = %d, want %d", matrix, got, want)
	}

	want = 1
	matrix = [][]byte{{0, 1}}
	if got := maximalRectangle(matrix); got != want {
		t.Errorf("maximalRectangle(%v) = %d, want %d", matrix, got, want)
	}
}
