package dp

import (
	"math"
)

// https://atcoder.jp/contests/dp/tasks/dp_a
func minPossibleTotalCost(n int, height []int) int {
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = 1 << 32
	}
	for i := 0; i < n; i++ {
		for _, j := range []int{i + 1, i + 2} {
			if j < n {
				cost := int(math.Abs(float64(height[i] - height[j])))
				dp[j] = min(dp[j], cost+dp[i])
			}
		}
	}
	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
