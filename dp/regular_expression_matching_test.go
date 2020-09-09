package dp

import (
	"testing"
)

func TestIsMatchRecursive(t *testing.T) {
	s, p := "aa", "a*"
	want := true
	if got := isMatch(s, p); got != want {
		t.Errorf("isMatchRecursive('%s', '%s') = %v, want %v", s, p, got, want)
	}
	s, p = "aa", "a"
	want = false
	if got := isMatch(s, p); got != want {
		t.Errorf("isMatchRecursive('%s', '%s') = %v, want %v", s, p, got, want)
	}
	s, p = "ab", ".*"
	want = true
	if got := isMatch(s, p); got != want {
		t.Errorf("isMatchRecursive('%s', '%s') = %v, want %v", s, p, got, want)
	}
	s, p = "aab", "c*a*b"
	want = true
	if got := isMatch(s, p); got != want {
		t.Errorf("isMatchRecursive('%s', '%s') = %v, want %v", s, p, got, want)
	}
	s, p = "mississippi", "mis*is*p*."
	want = false
	if got := isMatch(s, p); got != want {
		t.Errorf("isMatchRecursive('%s', '%s') = %v, want %v", s, p, got, want)
	}
}
