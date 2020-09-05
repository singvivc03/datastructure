package dp

// MaxCutNaive func
func MaxCutNaive(n, a, b, c int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	res := max2(MaxCutNaive(n-a, a, b, c), MaxCutNaive(n-b, a, b, c),
		MaxCutNaive(n-c, a, b, c))
	if res == -1 {
		return -1
	}
	return res + 1
}

// MaxCutDp func
func MaxCutDp(n, a, b, c int) int {
	memo := make([]int, n+1)
	memo[0] = 0
	cuts := []int{a, b, c}
	for i := 1; i <= n; i++ {
		memo[i] = -1
		for j := 0; j < len(cuts); j++ {
			if i-cuts[j] >= 0 {
				memo[i] = max(memo[i], memo[i-cuts[j]])
			}
		}
		if memo[i] != -1 {
			memo[i]++
		}
	}
	return memo[n]
}

func max2(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	return c
}
