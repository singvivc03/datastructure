package array

import (
	"testing"
)

func TestGenerateParanthesis(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	n := 3
	if got := generateParenthesis(3); !checkIfAllParanthesisIsGenerated(got, want) {
		t.Errorf("generateParanthesis(%d)=%v, want %v", n, got, want)
	}
}

func TestGenerateParanthesisSecondApproach(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	n := 3
	if got := generateParenthesisSecondApproach(3); !checkIfAllParanthesisIsGenerated(got, want) {
		t.Errorf("generateParenthesisSecondApproach(%d)=%v, want %v", n, got, want)
	}
}

func checkIfAllParanthesisIsGenerated(got []string, want []string) bool {
	paranthesis := make(map[string]bool)
	for _, element := range got {
		paranthesis[element] = true
	}
	for _, expected := range want {
		if !paranthesis[expected] {
			return false
		}
	}
	return true
}
