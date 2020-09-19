package dp

import (
	"testing"
)

func TestLongestPath(t *testing.T) {
	want := 3
	n, m := 5, 8
	graph := make(map[int][]int)
	from, to := []int{5, 2, 2, 5, 5, 1, 4, 1}, []int{3, 3, 4, 2, 1, 4, 3, 3}
	for i := 0; i < m; i++ {
		targets, _ := graph[from[i]]
		targets = append(targets, to[i])
		graph[from[i]] = targets
	}
	if got := longestPath(n, graph); got != want {
		t.Errorf("longestPath(%d, %v)=%d, want %d", n, graph, got, want)
	}

	want = 2
	n, m = 6, 3
	graph = make(map[int][]int)
	from, to = []int{2, 4, 5}, []int{3, 5, 6}
	for i := 0; i < m; i++ {
		targets, _ := graph[from[i]]
		targets = append(targets, to[i])
		graph[from[i]] = targets
	}
	if got := longestPath(n, graph); got != want {
		t.Errorf("longestPath(%d, %v)=%d, want %d", n, graph, got, want)
	}

	want = 3
	n, m = 4, 5
	graph = make(map[int][]int)
	from, to = []int{1, 1, 3, 2, 3}, []int{2, 3, 2, 4, 4}
	for i := 0; i < m; i++ {
		targets, _ := graph[from[i]]
		targets = append(targets, to[i])
		graph[from[i]] = targets
	}
	if got := longestPath(n, graph); got != want {
		t.Errorf("longestPath(%d, %v)=%d, want %d", n, graph, got, want)
	}
}
