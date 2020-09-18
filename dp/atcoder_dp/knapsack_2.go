package dp

import (
	"math"
)

var dp []int64

/**func main() {
	var n, w int64
	fmt.Scanf("%d%d", &n, &w)
	weight := make([]int64, n)
	values := make([]int64, n)
	for i := 0; i < int(n); i++ {
		fmt.Scanf("%d%d", &weight[i], &values[i])
	}
	ans := maximumPossibleSumKnapsack2(n, w, values, weight)
	fmt.Println(ans)
}**/

func createDPTable(v int64) []int64 {
	dp = make([]int64, v+1)
	for i := range dp {
		dp[i] = 1 << 32
	}
	dp[0] = 0
	return dp
}

func maximumPossibleSumKnapsack2(n int64, w int64, values []int64, weight []int64) int64 {
	var totalValues int64
	for i := range values {
		totalValues += values[i]
	}
	dp = createDPTable(totalValues)
	for items := 0; items < int(n); items++ {
		for value := totalValues - values[items]; value >= 0; value-- {
			dp[value+values[items]] = int64(math.Min(float64(dp[value+values[items]]), float64(dp[value]+weight[items])))
		}
	}
	var ans int64 = 0
	for i := 0; i <= int(totalValues); i++ {
		if dp[i] <= w {
			ans = int64(math.Max(float64(ans), float64(i)))
		}
	}
	return ans
}
