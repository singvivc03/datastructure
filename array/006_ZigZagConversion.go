package array

func convert(s string, numRows int) string {
	words := make([][]rune, numRows)
	pos, word := 0, 0
	increase := true
	for pos < len(s) {
		words[word] = append(words[word], rune(s[pos]))
		pos++
		if word == numRows-1 {
			increase = false
		}
		if word == 0 {
			increase = true
		}
		if increase {
			if numRows-1 != 0 {
				word++
			}
		} else {
			word--
		}
	}
	ans := ""
	for i := 0; i < numRows; i++ {
		ans += string(words[i])
	}
	return ans
}
