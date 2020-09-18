package dp

import (
	"testing"
)

func TestMaximumPossibleSumKnapsack2(t *testing.T) {
	want := int64(90)
	n, w := int64(3), int64(8)
	values := []int64{30, 50, 60}
	weight := []int64{3, 4, 5}

	if got := maximumPossibleSumKnapsack2(n, w, values, weight); got != want {
		t.Errorf("maximumPossibleSumKnapsack2(%d, %d, %v, %v)=%d, want %d", n, w, values, weight, got, want)
	}

	want = int64(17)
	n, w = int64(6), int64(15)
	values = []int64{5, 6, 4, 6, 5, 2}
	weight = []int64{6, 5, 6, 6, 3, 7}

	if got := maximumPossibleSumKnapsack2(n, w, values, weight); got != want {
		t.Errorf("maximumPossibleSumKnapsack2(%d, %d, %v, %v)=%d, want %d", n, w, values, weight, got, want)
	}
}
