package dp

// LongestPalindromicSubstring func finds the longest palindromic substring in a string
func LongestPalindromicSubstring(s string) string {
	table := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		table[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		table[i][i] = true
	}
	maxLength, startIndex := 1, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			startIndex = i
			table[i][i+1] = true
			maxLength = 2
		}
	}
	for i := 3; i <= len(s); i++ {
		for j := 0; j < len(s)-i+1; j++ {
			k := j + i - 1
			if table[j+1][k-1] && s[j] == s[k] {
				table[j][k] = true
				maxLength = i
				startIndex = j
			}
		}
	}
	return s[startIndex : maxLength+startIndex]
}
