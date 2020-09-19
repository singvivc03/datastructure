package dp

import (
	"math"
)

// https://atcoder.jp/contests/dp/tasks/dp_g
func longestPath(nodes int, graph map[int][]int) int {
	var maxPath int = -1 << 32
	visited := make([]bool, nodes+1)
	dp := make([]int, nodes+1)
	for i := 1; i <= nodes; i++ {
		_, ok := graph[i]
		if !visited[i] && ok {
			path := dfs(i, graph, dp, visited)
			maxPath = int(math.Max(float64(maxPath), float64(path)))
		}
	}
	return maxPath
}

func dfs(source int, graph map[int][]int, dp []int, visited []bool) int {
	edges := graph[source]
	if visited[source] {
		return dp[source]
	}
	visited[source] = true
	for i := 0; i < len(edges); i++ {
		path := 1 + dfs(edges[i], graph, dp, visited)
		dp[source] = int(math.Max(float64(dp[source]), float64(path)))
	}
	return dp[source]
}
