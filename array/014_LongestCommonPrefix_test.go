package array

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	want := "flow"
	strs := []string{"flower", "flow", "flowed"}
	if got := longestCommonPrefix(strs); got != want {
		t.Errorf("longestCommonPrefix(%v) = %s, want %s", strs, got, want)
	}
}
