package dp

import (
	"testing"
)

func TestCoinChangeNaive(t *testing.T) {
	want := 4
	if got := CoinChangeNaive([]int{1, 2, 3}, 4); got != want {
		t.Errorf("CoinChangeNaive() = %d, want %d", got, want)
	}
	want = 5
	if got := CoinChangeNaive([]int{2, 5, 3, 6}, 10); got != want {
		t.Errorf("CoinChangeNaive() = %d, want %d", got, want)
	}
	want = 3
	if got := CoinChangeNaive([]int{8, 3, 1, 2}, 3); got != want {
		t.Errorf("CoinChangeNaive() = %d, want %d", got, want)
	}
	want = 6
	if got := CoinChangeNaive([]int{1, 3, 5, 7}, 8); got != want {
		t.Errorf("CoinChangeNaive() = %d, want %d", got, want)
	}
}

func TestCoinChange(t *testing.T) {
	want := 4
	if got := CoinChange([]int{1, 2, 3}, 4); got != want {
		t.Errorf("CoinChange() = %d, want %d", got, want)
	}
	want = 5
	if got := CoinChange([]int{2, 5, 3, 6}, 10); got != want {
		t.Errorf("CoinChange() = %d, want %d", got, want)
	}
	want = 3
	if got := CoinChange([]int{8, 3, 1, 2}, 3); got != want {
		t.Errorf("CoinChange() = %d, want %d", got, want)
	}
	want = 6
	if got := CoinChangeNaive([]int{1, 3, 5, 7}, 8); got != want {
		t.Errorf("CoinChangeNaive() = %d, want %d", got, want)
	}
}
