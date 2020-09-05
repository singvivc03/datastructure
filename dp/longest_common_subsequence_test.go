package dp

import "testing"

func TestLongestCommonSubsequenceNaiveUsingBackwardIteration(t *testing.T) {
	want := 2
	s1, s2 := "AXYZ", "BAZ"
	if got := LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("LongestCommonSubsequenceNaiveUsingBackwardIteration(%s, %s, %d, %d) = %d, want %d", s1, s2, len(s1),
			len(s2), got, want)
	}

	want = 4
	s1, s2 = "ABXYZCD", "ABCD"
	if got := LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2, len(s1), len(s2)); got != want {
		t.Errorf("LongestCommonSubsequenceNaiveUsingBackwardIteration(%s, %s, %d, %d) = %d, want %d", s1, s2, len(s1),
			len(s2), got, want)
	}
}

func TestLongestCommonSubsequenceNaiveUsingFrontIteration(t *testing.T) {
	want := 2
	s1, s2 := "AXYZ", "BAZ"
	if got := LongestCommonSubsequenceNaiveUsingFrontIteration(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceNaiveUsingFrontIteration(%s, %s) = %d, want %d", s1, s2, got, want)
	}

	want = 4
	s1, s2 = "ABXYZCD", "ABCD"
	if got := LongestCommonSubsequenceNaiveUsingFrontIteration(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceNaiveUsingFrontIteration(%s, %s) = %d, want %d", s1, s2, got, want)
	}
}

func TestLongestCommonSubsequenceUsingDp(t *testing.T) {
	want := 2
	s1, s2 := "AXYZ", "BAZ"
	if got := LongestCommonSubsequenceUsingDp(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceUsingDp(%s, %s) = %d, want %d", s1, s2, got, want)
	}

	want = 4
	s1, s2 = "ABXYZCD", "ABCD"
	if got := LongestCommonSubsequenceUsingDp(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceUsingDp(%s, %s) = %d, want %d", s1, s2, got, want)
	}
}

func TestLongestCommonSubsequenceUsingDpTabulation(t *testing.T) {
	want := 2
	s1, s2 := "AXYZ", "BAZ"
	if got := LongestCommonSubsequenceUsingDpTabulation(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceUsingDpTabulation(%s, %s) = %d, want %d", s1, s2, got, want)
	}

	want = 4
	s1, s2 = "ABXYZCD", "ABCD"
	if got := LongestCommonSubsequenceUsingDpTabulation(s1, s2); got != want {
		t.Errorf("LongestCommonSubsequenceUsingDpTabulation(%s, %s) = %d, want %d", s1, s2, got, want)
	}
}
