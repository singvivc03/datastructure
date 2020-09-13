package dp

func uniquePath(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	memo := buildMemo(m, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i][j-1] + memo[i-1][j]
		}
	}
	return memo[m-1][n-1]
}

func buildMemo(m int, n int) [][]int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	memo[0][0] = 0
	for i := 1; i < n; i++ {
		memo[0][i] = 1
	}
	for i := 1; i < m; i++ {
		memo[i][0] = 1
	}
	return memo
}
