package array

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var ans []string
	constructParanthesis(&ans, "", 0, 0, n)
	return ans
}

func constructParanthesis(ans *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*ans = append(*ans, current)
		return
	}
	if open < max {
		constructParanthesis(ans, current+"(", open+1, close, max)
	}
	if close < open {
		constructParanthesis(ans, current+")", open, close+1, max)
	}
}

func generateParenthesisSecondApproach(n int) []string {
	var ans []string
	if n == 0 {
		ans = append(ans, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generateParenthesisSecondApproach(c) {
				for _, right := range generateParenthesisSecondApproach(n - c - 1) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
	}
	return ans
}
