package dp

import (
	"testing"
)

func TestLongestValidParanthesis(t *testing.T) {
	want := 10
	if got := longestValidParanthesis("((())())())"); got != want {
		t.Errorf("longestValidParanthesis() = %d, want %d", got, want)
	}
	want = 2
	if got := longestValidParanthesis("(()"); got != want {
		t.Errorf("longestValidParanthesis() = %d, want %d", got, want)
	}
	want = 4
	if got := longestValidParanthesis(")()())"); got != want {
		t.Errorf("longestValidParanthesis() = %d, want %d", got, want)
	}
}
