package dp

func optimalGameRecur(array []int, start, end, sum int) int {
	if start+1 == end {
		return max(array[start], array[end])
	}
	return max((sum - optimalGameRecur(array, start+1, end, sum-array[start])),
		(sum - optimalGameRecur(array, start, end-1, sum-array[end])))
}

func sol2(array []int, start int, end int) int {
	if start+1 == end {
		return max(array[start], array[end])
	}
	return max(array[start]+min2(sol2(array, start+1, end-1), sol2(array, start+2, end)),
		array[end]+min2(sol2(array, start+1, end-1), sol2(array, start, end-2)))
}

func solDP(array []int, size int) int {
	memo := buildMemoTable(size, size)
	for i := 1; i < size-1; i++ {
		memo[i][i+1] = max(array[i], array[i+1])
	}
	for gap := 3; gap < size; gap += 2 {
		for j := 0; j+gap < size; j++ {
			k := j + gap
			memo[j][k] = max(array[j]+min2(memo[j+1][k-1], memo[j+2][k]),
				array[k]+min2(memo[j+1][k-1], memo[j][k-2]))
		}
	}
	return memo[0][size-1]
}

func sol(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return optimalGameRecur(array, 0, len(array)-1, sum)
}
