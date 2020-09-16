package dp

import (
	"testing"
)

func TestMinPossibleTotalCost(t *testing.T) {
	want := 30
	n := 4
	height := []int{10, 30, 40, 20}
	if got := minPossibleTotalCost(n, height); got != want {
		t.Errorf("minPossibleTotalCost(%d, %v) = %d, want %d", n, height, got, want)
	}
	want = 0
	n = 2
	height = []int{10, 10}
	if got := minPossibleTotalCost(n, height); got != want {
		t.Errorf("minPossibleTotalCost(%d, %v) = %d, want %d", n, height, got, want)
	}
	want = 40
	n = 6
	height = []int{30, 10, 60, 10, 60, 50}
	if got := minPossibleTotalCost(n, height); got != want {
		t.Errorf("minPossibleTotalCost(%d, %v) = %d, want %d", n, height, got, want)
	}
}
