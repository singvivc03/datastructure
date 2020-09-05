package dp

// LongestIncreasingSubsequence func finds the longest increasing subsequence
func LongestIncreasingSubsequence(array []int, n int) int {
	lis := make([]int, n)
	lis[0] = 1
	ans := 0
	for i := 1; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if array[i] > array[j] {
				lis[i] = max(lis[i], lis[j]+1)
				ans = max(ans, lis[i])
			}
		}
	}
	return ans
}

// FastLongestIncreasingSubsequence func
func FastLongestIncreasingSubsequence(array []int, n int) int {
	tail := make([]int, n)
	tail[0] = array[0]
	len := 1
	for i := 1; i < n; i++ {
		if array[i] > tail[len-1] {
			tail[i] = array[i]
			len++
		} else {
			idx := searchSmallerThan(tail, 0, len-1, array[i])
			tail[idx] = array[i]
		}
	}
	return len
}

func searchSmallerThan(tail []int, start, end, target int) int {
	for start < end {
		mid := (start + end) / 2
		if tail[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
