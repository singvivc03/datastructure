package dp

import (
	"math"
)

func minJumpRecursive(array []int, pos int) int {
	if pos == 1 {
		return 0
	}
	res := math.MaxInt64
	for i := 0; i < pos-1; i++ {
		if i+array[i] >= pos-1 {
			subRes := minJumpRecursive(array, i+1)
			if subRes != math.MaxInt64 {
				res = min2(res, subRes+1)
			}
		}
	}
	return res
}

func minJumpDP(array []int, pos int) int {
	memo := make([]int, pos)
	memo[0] = 0
	for i := 1; i < pos; i++ {
		memo[i] = math.MaxInt64
	}
	for i := 1; i < pos; i++ {
		for j := 0; j < i; j++ {
			if array[j]+j >= i {
				if memo[j] != math.MaxInt64 {
					memo[i] = min2(memo[i], memo[j]+1)
				}
			}
		}
	}
	return memo[pos-1]
}
