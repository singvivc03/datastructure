package dp

import (
	"testing"
)

func TestMinJumpRecursive(t *testing.T) {
	want := 2
	array := []int{3, 4, 2, 1, 2, 1}
	if got := minJumpRecursive(array, len(array)); got != want {
		t.Errorf("minJumpRecursive() = %d, want %d", got, want)
	}
}

func TestMinJumpDP(t *testing.T) {
	want := 2
	array := []int{3, 4, 2, 1, 2, 1}
	if got := minJumpDP(array, len(array)); got != want {
		t.Errorf("minJumpDP() = %d, want %d", got, want)
	}
}
