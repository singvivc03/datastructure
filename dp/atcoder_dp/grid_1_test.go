package dp

import (
	"testing"
)

func TestNumberOfPath(t *testing.T) {
	want := 3
	grid := [][]rune{{'.', '.', '.', '#'}, {'.', '#', '.', '.'}, {'.', '.', '.', '.'}}
	h, w := 3, 4
	if got := numberOfPath(h, w, grid); got != want {
		t.Errorf("numberOfPath(%d, %d, %v) = %d, want %d", h, w, grid, got, want)
	}

	want = 0
	grid = [][]rune{{'.', '.'}, {'#', '.'}, {'.', '.'}, {'.', '#'}, {'.', '.'}}
	h, w = 5, 2
	if got := numberOfPath(h, w, grid); got != want {
		t.Errorf("numberOfPath(%d, %d, %v) = %d, want %d", h, w, grid, got, want)
	}

	want = 24
	grid = [][]rune{{'.', '.', '#', '.', '.'}, {'.', '.', '.', '.', '.'}, {'#', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.'}, {'.', '.', '#', '.', '.'}}
	h, w = 5, 5
	if got := numberOfPath(h, w, grid); got != want {
		t.Errorf("numberOfPath(%d, %d, %v) = %d, want %d", h, w, grid, got, want)
	}
}
