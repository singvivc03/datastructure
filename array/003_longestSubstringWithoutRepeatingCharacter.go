package array

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	start, length := 0, -1
	alreadySeen := make(map[rune]int)
	for end := 0; end < len(s); end++ {
		current := rune(s[end])
		if _, ok := alreadySeen[current]; ok {
			if start <= alreadySeen[current] {
				start = alreadySeen[current] + 1
			}
		}
		length = int(math.Max(float64(length), float64(end-start+1)))
		alreadySeen[current] = end
	}
	return length
}
