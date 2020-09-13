package dp

func minDistance(word1 string, word2 string) int {
	return minDistanceRecursive(word1, word2, len(word1), len(word2))
}

func minDistanceRecursive(word1 string, word2 string, m int, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if word1[m-1] == word2[n-1] {
		return minDistanceRecursive(word1, word2, m-1, n-1)
	}
	return 1 + min(minDistanceRecursive(word1, word2, m, n-1),
		minDistanceRecursive(word1, word2, m-1, n),
		minDistanceRecursive(word1, word2, m-1, n-1))
}

func minDistanceDP(word1 string, word2 string) int {
	memo := buildMemoTable(len(word1)+1, len(word2)+1)
	for i := 1; i <= len(word2); i++ {
		memo[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		memo[i][0] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = 1 + min(memo[i][j-1], memo[i-1][j], memo[i-1][j-1])
			}
		}
	}
	return memo[len(word1)][len(word2)]
}
