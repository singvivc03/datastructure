package dp

import (
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	want := 4
	if got := LongestIncreasingSubsequence([]int{3, 4, 2, 8, 10, 5, 1}, 7); got != want {
		t.Errorf("LongestIncreasingSubsequence() = %d, want %d", got, want)
	}
}

func TestFastLongestIncreasingSubsequence(t *testing.T) {
	want := 4
	if got := FastLongestIncreasingSubsequence([]int{3, 4, 2, 8, 10, 5, 1}, 7); got != want {
		t.Errorf("FastLongestIncreasingSubsequence() = %d, want %d", got, want)
	}
}
