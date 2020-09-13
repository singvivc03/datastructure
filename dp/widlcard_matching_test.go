package dp

import (
	"testing"
)

func TestIsWildCardMatch(t *testing.T) {
	want := true
	if got := isWildCardMatch("adceb", "*a*b"); got != want {
		t.Errorf("isWildCardMatch('adceb', '*a*b') = %v, want %v", got, want)
	}
	want = false
	if got := isWildCardMatch("aa", "a"); got != want {
		t.Errorf("isWildCardMatch('aa', 'a') = %v, want %v", got, want)
	}
	want = false
	if got := isWildCardMatch("cb", "?a"); got != want {
		t.Errorf("isWildCardMatch('cb', '?a') = %v, want %v", got, want)
	}
	want = true
	if got := isWildCardMatch("aa", "*"); got != want {
		t.Errorf("isWildCardMatch('aa', '*') = %v, want %v", got, want)
	}
	want = false
	if got := isWildCardMatch("acdcb", "a*c?b"); got != want {
		t.Errorf("isWildCardMatch('acdcb', 'a*c?b') = %v, want %v", got, want)
	}
}
