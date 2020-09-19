package dp

import (
	"testing"
)

func TestLcs(t *testing.T) {
	want := "aaadara"
	str1, str2 := "abracadabra", "avadakedavra"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}
	want = ""
	str1, str2 = "a", "z"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = "gtab"
	str1, str2 = "aggtab", "gxtxayb"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = ""
	str1, str2 = "aggtab", ""
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = ""
	str1, str2 = "", "aggtab"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = "ay"
	str1, str2 = "axy", "ayz"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = "axb"
	str1, str2 = "axyb", "abyxb"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}

	want = "aa"
	str1, str2 = "aa", "xayaz"
	if got := lcs(str1, str2); got != want {
		t.Errorf("lcs(%s, %s) = %s, want %s", str1, str2, got, want)
	}
}
