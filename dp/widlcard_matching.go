package dp

func isWildCardMatch(s string, p string) bool {
	memo := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		memo[i] = make([]bool, len(p)+1)
	}
	memo[0][0] = true
	for i := 1; i <= len(p); i++ {
		memo[0][i] = false
		if p[i-1] == '*' {
			memo[0][i] = memo[0][i-1]
		}
	}
	return isWildCardAMatch(s, p, memo)
}

func isWildCardAMatch(s string, p string, memo [][]bool) bool {
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				memo[i][j] = memo[i-1][j-1]
			} else if p[j-1] == '*' {
				memo[i][j] = memo[i-1][j] || memo[i][j-1]
			}
		}
	}
	return memo[len(s)][len(p)]
}
