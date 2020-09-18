package dp

import (
	"testing"
)

func TestMaximumPossibleSum(t *testing.T) {
	want := int64(90)
	n, w := int64(3), int64(8)
	weight := []int64{3, 4, 5}
	values := []int64{30, 50, 60}
	if got := maximumPossibleSum(n, w, values, weight); got != want {
		t.Errorf("maximumPossibleSum(%d, %d, %v, %v)=%d, want %d", n, w, values, weight, got, want)
	}

	want = int64(5000000000)
	n, w = int64(5), int64(5)
	weight = []int64{1, 1, 1, 1, 1}
	values = []int64{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	if got := maximumPossibleSum(n, w, values, weight); got != want {
		t.Errorf("maximumPossibleSum(%d, %d, %v, %v)=%d, want %d", n, w, values, weight, got, want)
	}

	want = int64(10)
	n, w = int64(1), int64(1000000000)
	weight = []int64{1000000000}
	values = []int64{10}
	if got := maximumPossibleSum(n, w, values, weight); got != want {
		t.Errorf("maximumPossibleSum(%d, %d, %v, %v)=%d, want %d", n, w, values, weight, got, want)
	}
}
