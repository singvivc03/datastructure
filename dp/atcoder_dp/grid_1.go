package dp

import (
	"math"
)

// MOD ...
var MOD int64 = int64(math.Pow(float64(10), float64(9)) + 7)

/**func main() {
	var h, w int
	fmt.Scanf("%d%d", &h, &w)
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]rune, w)
		var line string
		fmt.Scanf("%s", &line)
		for j := 0; j < w; j++ {
			grid[i][j] = rune(line[j])
		}
	}
	ans := numberOfPath(h, w, grid)
	fmt.Println(ans)
}**/

// https://atcoder.jp/contests/dp/tasks/dp_h
func numberOfPath(h int, w int, grid [][]rune) int {
	dp := createTable(h-1, w-1)
	for i := 0; i < w; i++ {
		if grid[0][i] == '#' {
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < h; i++ {
		if grid[i][0] == '#' {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if grid[i][j] == '.' {
				dp[i][j] = int(int64(dp[i-1][j]+dp[i][j-1]) % MOD)
			}
		}
	}
	return dp[h-1][w-1]
}
