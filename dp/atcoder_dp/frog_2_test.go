package dp

import (
	"testing"
)

func TestMinPossibleTotalCost2(t *testing.T) {
	want := 30
	n := 5
	k := 3
	height := []int{10, 30, 40, 50, 20}
	if got := minPossibleTotalCost2(n, k, height); got != want {
		t.Errorf("minPossibleTotalCost2(%d, %d, %v) = %d, want %d", n, k, height, got, want)
	}
	want = 20
	n = 3
	k = 1
	height = []int{10, 20, 10}
	if got := minPossibleTotalCost2(n, k, height); got != want {
		t.Errorf("minPossibleTotalCost2(%d, %d, %v) = %d, want %d", n, k, height, got, want)
	}
	want = 0
	n = 2
	k = 100
	height = []int{10, 10}
	if got := minPossibleTotalCost2(n, k, height); got != want {
		t.Errorf("minPossibleTotalCost2(%d, %d, %v) = %d, want %d", n, k, height, got, want)
	}
	want = 40
	n = 10
	k = 4
	height = []int{40, 10, 20, 70, 80, 10, 20, 70, 80, 60}
	if got := minPossibleTotalCost2(n, k, height); got != want {
		t.Errorf("minPossibleTotalCost2(%d, %d, %v) = %d, want %d", n, k, height, got, want)
	}
}
