package dp

import (
	"testing"
)

func TestMinCoin(t *testing.T) {
	want := 2
	if got := minCoin([]int{25, 10, 5}, 30); got != want {
		t.Errorf("minCoin() = %d, want %d", got, want)
	}
}

func TestMinCoinDp(t *testing.T) {
	want := 2
	if got := minCoinDp([]int{3, 4, 1}, 5); got != want {
		t.Errorf("minCoinDp() = %d, want %d", got, want)
	}
	want = 2
	if got := minCoinDp([]int{25, 10, 5}, 30); got != want {
		t.Errorf("minCoinDp() = %d, want %d", got, want)
	}
}
