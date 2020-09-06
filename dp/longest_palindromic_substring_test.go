package dp

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	want := "ccc"
	if got := LongestPalindromicSubstring("ccc"); got != want {
		t.Errorf("LongestPalindromicSubstring(babad) = %s, want %s", got, want)
	}
	want = "aba"
	if got := LongestPalindromicSubstring("babad"); got != want {
		t.Errorf("LongestPalindromicSubstring(babad) = %s, want %s", got, want)
	}
	want = "aabaa"
	if got := LongestPalindromicSubstring("vivekaabaa"); got != want {
		t.Errorf("LongestPalindromicSubstring(babad) = %s, want %s", got, want)
	}

	want = "geeksskeeg"
	s := "forgeeksskeegfor"
	if got := LongestPalindromicSubstring(s); got != want {
		t.Errorf("LongestPalindromicSubstring('%s') = %s, want %s", s, got, want)
	}
}
