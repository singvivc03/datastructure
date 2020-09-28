package array

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack, stackLength := make([]rune, len(s)), 0
	stack[0] = rune(s[0])
	stackLength++
	for i := 1; i < len(s); i++ {
		current := rune(s[i])
		if stackLength != 0 {
			previous := stack[stackLength-1]
			if (previous == '(' && current == ')') || (previous == '[' && current == ']') ||
				(previous == '{' && current == '}') {
				stackLength--
			} else {
				stack[stackLength] = rune(s[i])
				stackLength++
			}
		} else {
			stack[stackLength] = rune(s[i])
			stackLength++
		}
	}
	return stackLength == 0
}
