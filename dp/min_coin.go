package dp

import (
	"math"
)

func minCoin(coins []int, val int) int {
	if val == 0 {
		return 0
	}
	res := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		if coins[i] <= val {
			subRes := minCoin(coins, (val - coins[i]))
			if subRes != math.MaxInt64 {
				res = min2(res, subRes+1)
			}
		}
	}
	return res
}

func minCoinDp(coins []int, val int) int {
	memo := make([]int, val+1)
	memo[0] = 0
	for i := 1; i <= val; i++ {
		memo[i] = math.MaxInt64
	}
	for i := 1; i <= val; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				res := memo[i-coins[j]]
				if res != math.MaxInt64 {
					memo[i] = min2(memo[i], res+1)
				}
			}
		}
	}
	return memo[val]
}

func min2(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
