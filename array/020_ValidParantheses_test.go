package array

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	want := true
	s := "()"
	if got := isValid(s); got != want {
		t.Errorf("isValid(%s)=%v, want %v", s, got, want)
	}

	want = true
	s = "([{}])"
	if got := isValid(s); got != want {
		t.Errorf("isValid(%s)=%v, want %v", s, got, want)
	}

	want = true
	s = "()((([[{{{}}}]])))"
	if got := isValid(s); got != want {
		t.Errorf("isValid(%s)=%v, want %v", s, got, want)
	}

	want = false
	s = "(]"
	if got := isValid(s); got != want {
		t.Errorf("isValid(%s)=%v, want %v", s, got, want)
	}
}
