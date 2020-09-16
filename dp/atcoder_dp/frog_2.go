package dp

import (
	"math"
)

// https://atcoder.jp/contests/dp/tasks/dp_b
func minPossibleTotalCost2(n int, k int, height []int) int {
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = 1 << 32
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j <= k+i; j++ {
			if j < n {
				cost := int(math.Abs(float64(height[i] - height[j])))
				dp[j] = min(dp[j], dp[i]+cost)
			}
		}
	}
	return dp[n-1]
}
