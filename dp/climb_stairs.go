package dp

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsDP(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	for i := 1; i <= n; i++ {
		if i == 1 {
			memo[i] = memo[i-1]
		} else {
			memo[i] = memo[i-1] + memo[i-2]
		}
	}
	return memo[n]
}
