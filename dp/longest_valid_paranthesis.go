package dp

func longestValidParanthesis(s string) int {
	parenthesesStack, indexStack := make([]byte, len(s)+1), make([]int, len(s)+1)
	pindex, iindex, maxLength := 1, 1, 0
	indexStack[0] = -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			parenthesesStack[pindex] = s[i]
			pindex++
			indexStack[iindex] = i
			iindex++
		} else if s[i] == ')' && parenthesesStack[pindex-1] == '(' {
			pindex--
			iindex--
			last := indexStack[iindex-1]
			maxLength = max(maxLength, i-last)
		} else {
			indexStack[iindex] = i
			iindex++
		}
	}
	return maxLength
}
