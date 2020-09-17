package dp

import (
	"testing"
)

func TestMaximumTotalPoints(t *testing.T) {
	want := 210
	n := 3
	activities := [][]int{{10, 40, 70}, {20, 50, 80}, {30, 60, 90}}
	if got := maximumTotalPoints(n, activities); got != want {
		t.Errorf("maximumTotalPoints(%d, %v) = %d, want %d", n, activities, got, want)
	}
	want = 100
	n = 1
	activities = [][]int{{100, 10, 1}}
	if got := maximumTotalPoints(n, activities); got != want {
		t.Errorf("maximumTotalPoints(%d, %v) = %d, want %d", n, activities, got, want)
	}
	want = 46
	n = 7
	activities = [][]int{{6, 7, 8}, {8, 8, 3}, {2, 5, 2}, {7, 8, 6}, {4, 6, 8}, {2, 3, 4}, {7, 5, 1}}
	if got := maximumTotalPoints(n, activities); got != want {
		t.Errorf("maximumTotalPoints(%d, %v) = %d, want %d", n, activities, got, want)
	}
}
