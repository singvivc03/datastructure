package array

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	return commonPrefix(strs)
}

func commonPrefix(strs []string) string {
	lengthOfLongestSubstring := findLengthOfSmallestWord(strs)
	low, high := 0, lengthOfLongestSubstring
	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return strs[0][0 : (low+high)/2]
}

func findLengthOfSmallestWord(strs []string) int {
	minLength := 1 << 32
	for i := 0; i < len(strs); i++ {
		minLength = int(math.Min(float64(len(strs[i])), float64(minLength)))
	}
	return minLength
}

func isCommonPrefix(strs []string, length int) bool {
	partial := strs[0][0:length]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], partial) {
			return false
		}
	}
	return true
}
