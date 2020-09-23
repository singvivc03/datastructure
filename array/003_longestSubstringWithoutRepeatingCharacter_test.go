package array

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	want := 8
	s := "aaabbbcdgexya"
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("lengthOfLongestSubstring(%s)=%d, want %d", s, got, want)
	}
	want = 1
	s = "bbbbb"
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("lengthOfLongestSubstring(%s)=%d, want %d", s, got, want)
	}
	want = 3
	s = "pwwkew"
	if got := lengthOfLongestSubstring(s); got != want {
		t.Errorf("lengthOfLongestSubstring(%s)=%d, want %d", s, got, want)
	}
}
