package dp

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	want := 7
	if got := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}); got != want {
		t.Errorf("minPathSum() = %d, want %d", got, want)
	}
	want = 1
	if got := minPathSum([][]int{{1}}); got != want {
		t.Errorf("minPathSum() = %d, want %d", got, want)
	}
	want = 0
	if got := minPathSum([][]int{}); got != want {
		t.Errorf("minPathSum() = %d, want %d", got, want)
	}
}
