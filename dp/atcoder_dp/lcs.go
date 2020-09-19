package dp

import (
	"math"
)

func lcs(str1 string, str2 string) string {
	dp := createTable(len(str1), len(str2))
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	var ans []rune
	j := len(str2)
	for i := len(str1); i > 0 && j > 0; {
		if dp[i][j] == int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))) {
			if dp[i-1][j] >= dp[i][j-1] {
				i--
			} else {
				j--
			}
		} else if dp[i][j] == dp[i-1][j-1]+1 {
			ans = append(ans, rune(str1[i-1]))
			i--
			j--
		}
	}
	return reverse(string(ans))
}

func reverse(str string) string {
	if len(str) == 0 {
		return ""
	}
	reversed := make([]rune, len(str))
	for i := 0; i <= len(str)/2; i++ {
		reversed[i], reversed[len(str)-1-i] = rune(str[len(str)-i-1]), rune(str[i])
	}
	return string(reversed)
}

func createTable(n int, m int) [][]int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	return dp
}
