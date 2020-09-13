package dp

import (
	"testing"
)

func TestUniquePath(t *testing.T) {
	want := 28
	if got := uniquePath(3, 7); got != want {
		t.Errorf("uniquePath(%d, %d) = %d, want %d", 3, 7, got, want)
	}

	want = 3
	if got := uniquePath(3, 2); got != want {
		t.Errorf("uniquePath(%d, %d) = %d, want %d", 3, 2, got, want)
	}
	want = 0
	if got := uniquePath(0, 0); got != want {
		t.Errorf("uniquePath(%d, %d) = %d, want %d", 0, 0, got, want)
	}
}
