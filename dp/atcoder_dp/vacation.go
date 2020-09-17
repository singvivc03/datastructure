package dp

import (
	"math"
)

/*func main() {
	var n int
	fmt.Scanf("%d", &n)
	activities := make([][]int, n)
	for i := 0; i < n; i++ {
		activities[i] = make([]int, 3)
		fmt.Scanf("%d%d%d", &activities[i][0], &activities[i][1], &activities[i][2])
	}
	ans := maximumTotalPoints(n, activities)
	fmt.Println(ans)
}*/

// https://atcoder.jp/contests/dp/tasks/dp_c
func maximumTotalPoints(n int, activities [][]int) int {
	dp := make([]int, 3)
	for i := 0; i < n; i++ {
		newDP := make([]int, 3)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j != k {
					newDP[k] = int(math.Max(float64(newDP[k]), float64(dp[j]+activities[i][j])))
				}
			}
		}
		dp = newDP
	}
	return int(math.Max(float64(dp[0]), math.Max(float64(dp[1]), float64(dp[2]))))
}
