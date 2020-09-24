package array

import (
	"testing"
)

func TestReverse(t *testing.T) {
	want := 321
	x := 123
	if got := reverse(x); got != want {
		t.Errorf("reverse(%d)=%d, want %d", x, got, want)
	}

	want = -321
	x = -123
	if got := reverse(x); got != want {
		t.Errorf("reverse(%d)=%d, want %d", x, got, want)
	}

	want = 21
	x = 120
	if got := reverse(x); got != want {
		t.Errorf("reverse(%d)=%d, want %d", x, got, want)
	}

	want = 0
	x = 1534236469
	if got := reverse(x); got != want {
		t.Errorf("reverse(%d)=%d, want %d", x, got, want)
	}
}
