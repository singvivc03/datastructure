package dp

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	want := 2
	if got := climbStairs(2); got != want {
		t.Errorf("climbStairs(%d) = %d, want %d", 2, got, want)
	}
	want = 3
	if got := climbStairs(3); got != want {
		t.Errorf("climbStairs(%d) = %d, want %d", 2, got, want)
	}
	want = 8
	if got := climbStairs(5); got != want {
		t.Errorf("climbStairs(%d) = %d, want %d", 2, got, want)
	}
}

func TestClimbStairsDP(t *testing.T) {
	want := 2
	if got := climbStairsDP(2); got != want {
		t.Errorf("climbStairsDP(%d) = %d, want %d", 2, got, want)
	}
	want = 3
	if got := climbStairsDP(3); got != want {
		t.Errorf("climbStairsDP(%d) = %d, want %d", 2, got, want)
	}
	want = 8
	if got := climbStairsDP(5); got != want {
		t.Errorf("climbStairsDP(%d) = %d, want %d", 2, got, want)
	}
}
