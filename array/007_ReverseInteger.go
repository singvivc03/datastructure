package array

import (
	"math"
)

func reverse(x int) int {
	ans := 0
	min, max := math.MinInt32, math.MaxInt32
	for x/10 != 0 {
		ans += (x % 10)
		x = x / 10
		ans = ans * 10
	}
	ans += x
	if ans > max || ans < min {
		ans = 0
	}
	return ans
}
