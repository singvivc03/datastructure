package dp

import (
	"testing"
)

func TestOptimalGameRecur(t *testing.T) {
	want := 25
	if got := sol([]int{20, 5, 4, 6}); got != want {
		t.Errorf("sol() = %d, want %d", got, want)
	}
	want = 32
	if got := sol([]int{20, 5, 4, 6, 8, 3}); got != want {
		t.Errorf("sol() = %d, want %d", got, want)
	}
}

func TestOptimalGameDP(t *testing.T) {
	want := 25
	if got := solDP([]int{20, 5, 4, 6}, 4); got != want {
		t.Errorf("sol() = %d, want %d", got, want)
	}
	want = 32
	if got := solDP([]int{20, 5, 4, 6, 8, 3}, 6); got != want {
		t.Errorf("sol() = %d, want %d", got, want)
	}
}
