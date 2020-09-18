package dp

import (
	"math"
)

/*func main() {
	var n, w int64
	fmt.Scanf("%d%d", &n, &w)
	weight := make([]int64, n)
	values := make([]int64, n)
	for i := 0; i < int(n); i++ {
		fmt.Scanf("%d%d", &weight[i], &values[i])
	}
	ans := maximumPossibleSum(n, w, values, weight)
	fmt.Println(ans)
}*/

// https://atcoder.jp/contests/dp/tasks/dp_d
func maximumPossibleSum(n int64, w int64, values []int64, weight []int64) int64 {
	dp := make([][]int64, n+1)
	for i := 0; i <= int(n); i++ {
		dp[i] = make([]int64, w+1)
	}
	for i := 1; i <= int(n); i++ {
		for j := 1; j <= int(w); j++ {
			if weight[i-1] > int64(j) {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int64(math.Max(float64(dp[i-1][j]), float64(values[i-1]+dp[i-1][int64(j)-weight[i-1]])))
			}
		}
	}
	return dp[n][w]
}

func maximumPossibleSumRecur(n int64, w int64, values []int64, weight []int64) int64 {
	if n == 0 || w == 0 {
		return 0
	}
	if weight[n-1] > w {
		return maximumPossibleSumRecur(n-1, w, values, weight)
	}
	return int64(math.Max(float64(maximumPossibleSumRecur(n-1, w, values, weight)),
		float64(values[n-1]+maximumPossibleSumRecur(n-1, w-weight[n-1], values, weight))))
}
