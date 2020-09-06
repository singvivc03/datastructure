package dp

func knapsackRecur(weight []int, value []int, w int, n int) int {
	if n == 0 || w == 0 {
		return 0
	}
	if weight[n-1] > w {
		return knapsackRecur(weight, value, w, n-1)
	}
	return max(knapsackRecur(weight, value, w, n-1),
		value[n-1]+knapsackRecur(weight, value, w-weight[n-1], n-1))
}

func knapsackDp(weight []int, value []int, w int, n int) int {
	memo := buildMemoTable(n+1, w+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if weight[i-1] > j {
				memo[i][j] = memo[i-1][j]
			} else {
				memo[i][j] = max(memo[i-1][j], value[i-1]+memo[i-1][j-weight[i-1]])
			}
		}
	}
	return memo[n][w]
}
