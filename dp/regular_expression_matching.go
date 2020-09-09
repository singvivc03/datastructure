package dp

func isMatch(s string, p string) bool {
	memo := buildTable(p, len(s), len(p))
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				memo[i][j] = memo[i-1][j-1]
			} else if p[j-1] == '*' {
				memo[i][j] = memo[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					memo[i][j] = memo[i][j] || memo[i-1][j]
				}
			}
		}
	}
	return memo[len(s)][len(p)]
}

func buildTable(p string, n, m int) [][]bool {
	memo := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]bool, m+1)
	}
	memo[0][0] = true
	for i := 1; i <= m; i++ {
		if p[i-1] == '*' {
			memo[0][i] = memo[0][i-2]
		}
	}
	return memo
}
