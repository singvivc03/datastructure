package dp

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	want := 3
	word1, word2 := "horse", "ros"
	if got := minDistance(word1, word2); got != want {
		t.Errorf("minDistance(%s, %s) = %d, want %d", word1, word2, got, want)
	}
	want = 5
	word1, word2 = "intention", "execution"
	if got := minDistance(word1, word2); got != want {
		t.Errorf("minDistance(%s, %s) = %d, want %d", word1, word2, got, want)
	}
	want = 2
	word1, word2 = "ab", "bc"
	if got := minDistance(word1, word2); got != want {
		t.Errorf("minDistance(%s, %s) = %d, want %d", word1, word2, got, want)
	}
}

func TestMinDistanceDP(t *testing.T) {
	want := 3
	word1, word2 := "horse", "ros"
	if got := minDistanceDP(word1, word2); got != want {
		t.Errorf("minDistanceDP(%s, %s) = %d, want %d", word1, word2, got, want)
	}
	want = 5
	word1, word2 = "intention", "execution"
	if got := minDistanceDP(word1, word2); got != want {
		t.Errorf("minDistance(%s, %s) = %d, want %d", word1, word2, got, want)
	}
	want = 7
	word1, word2 = "dinitrophenylhydrazine", "benzalphenylhydrazone"
	if got := minDistanceDP(word1, word2); got != want {
		t.Errorf("minDistance(%s, %s) = %d, want %d", word1, word2, got, want)
	}
}
