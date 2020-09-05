package dp

// LongestCommonSubsequenceNaiveUsingBackwardIteration calculates the longest common subsequence of two strings using the naive
// approach
func LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if s1[m-1] == s2[n-1] {
		return 1 + LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2, m-1, n-1)
	}
	return max(LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2, m-1, n),
		LongestCommonSubsequenceNaiveUsingBackwardIteration(s1, s2, m, n-1))
}

// LongestCommonSubsequenceNaiveUsingFrontIteration calculates the longest common subsequence of two strings using the naive
// approach
func LongestCommonSubsequenceNaiveUsingFrontIteration(s1, s2 string) int {
	return longestCommonSubsequenceNaive(s1, s2, 0, 0)
}

func longestCommonSubsequenceNaive(s1, s2 string, m, n int) int {
	if m >= len(s1) || n >= len(s2) {
		return 0
	}
	if s1[m] == s2[n] {
		return 1 + longestCommonSubsequenceNaive(s1, s2, m+1, n+1)
	}
	return max(longestCommonSubsequenceNaive(s1, s2, m+1, n),
		longestCommonSubsequenceNaive(s1, s2, m, n+1))
}

func makeTable(n, m int) [][]int {
	memo := make([][]int, n)
	for index := range memo {
		memo[index] = make([]int, m)
		for pos := range memo[index] {
			memo[index][pos] = -1
		}
	}
	return memo
}

// LongestCommonSubsequenceUsingDp calculates the longest common subsequence of two strings using the dynamic programming
// approach (Memoization)
func LongestCommonSubsequenceUsingDp(s1, s2 string) int {
	memo := makeTable(len(s1)+1, len(s2)+1)
	return longestCommonSubsequenceDp(s1, s2, memo, len(s1), len(s2))
}

func longestCommonSubsequenceDp(s1, s2 string, memo [][]int, m int, n int) int {
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if m == 0 || n == 0 {
		memo[m][n] = 0
	} else {
		if s1[m-1] == s2[n-1] {
			memo[m][n] = 1 + longestCommonSubsequenceDp(s1, s2, memo, m-1, n-1)
		} else {
			memo[m][n] = max(longestCommonSubsequenceDp(s1, s2, memo, m-1, n),
				longestCommonSubsequenceDp(s1, s2, memo, m, n-1))
		}
	}
	return memo[m][n]
}

// LongestCommonSubsequenceUsingDpTabulation calculates the longest common subsequence of two strings using the dynamic programming
// approach (Tabulation)
func LongestCommonSubsequenceUsingDpTabulation(s1, s2 string) int {
	memo := makeTable(len(s1)+1, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		memo[i][0] = 0
	}
	for i := 0; i <= len(s2); i++ {
		memo[0][i] = 0
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}
	return memo[len(s1)][len(s2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
