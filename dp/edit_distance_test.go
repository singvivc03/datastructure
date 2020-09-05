package dp

import (
	"testing"
)

func TestEditDistanceNaive(t *testing.T) {
	want := 3
	s1, s2 := "Saturday", "Sunday"
	if got := EditDistanceNaive(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceNaive(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
	want = 1
	s1, s2 = "Cat", "Cut"
	if got := EditDistanceNaive(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceNaive(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
}

func TestEditDistanceRecurseDP(t *testing.T) {
	want := 3
	s1, s2 := "Saturday", "Sunday"
	if got := EditDistanceRecurseDP(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceRecurseDP(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
	want = 1
	s1, s2 = "Cat", "Cut"
	if got := EditDistanceRecurseDP(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceRecurseDP(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
}

func TestEditDistanceTabulation(t *testing.T) {
	want := 3
	s1, s2 := "Saturday", "Sunday"
	if got := EditDistanceTabulation(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceTabulation(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
	want = 1
	s1, s2 = "Cat", "Cut"
	if got := EditDistanceTabulation(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("\nEditDistanceTabulation(%v, %v, %d, %d) = %d, want %d", s1, s2, len(s1), len(s2), got, want)
	}
}
