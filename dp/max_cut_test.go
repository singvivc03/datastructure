package dp

import (
	"testing"
)

func TestMaxCutNaive(t *testing.T) {
	want := 5
	if got := MaxCutNaive(5, 1, 2, 3); want != got {
		t.Errorf("MaxCutNaive() = %d, want %d", got, want)
	}
}

func TestMaxCutDp(t *testing.T) {
	want := 5
	if got := MaxCutDp(5, 1, 2, 3); want != got {
		t.Errorf("MaxCutDp() = %d, want %d", got, want)
	}
	want = -1
	if got := MaxCutDp(5, 2, 2, 2); want != got {
		t.Errorf("MaxCutDp() = %d, want %d", got, want)
	}
}
