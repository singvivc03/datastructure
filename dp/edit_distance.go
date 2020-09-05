package dp

// EditDistanceNaive func naively solves the minimum number of operations required to convert s1 to s2 using
// insert, delete and replace
func EditDistanceNaive(s1, s2 string, m, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if s1[m-1] == s2[n-1] {
		return EditDistanceNaive(s1, s2, m-1, n-1)
	}
	return 1 + min(EditDistanceNaive(s1, s2, m-1, n), EditDistanceNaive(s1, s2, m-1, n-1),
		EditDistanceNaive(s1, s2, m, n-1))
}

// EditDistanceRecurseDP func solves the minimum number of operations required to convert s1 to s2 using
// insert, delete and replace, recursively using dynamic programing
func EditDistanceRecurseDP(s1, s2 string, m, n int) int {
	memo := buildMemoTable(m+1, n+1)
	memo[0][0] = 0
	for i := 1; i <= n; i++ {
		memo[0][i] = i
	}
	for i := 1; i <= m; i++ {
		memo[i][0] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	return editDistanceRecurse(s1, s2, memo, m, n)
}

func editDistanceRecurse(s1, s2 string, memo [][]int, m, n int) int {
	if m == 0 {
		return memo[0][n]
	}
	if n == 0 {
		return memo[m][0]
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if s1[m-1] == s2[n-1] {
		memo[m][n] = editDistanceRecurse(s1, s2, memo, m-1, n-1)
	} else {
		memo[m][n] = 1 + min(editDistanceRecurse(s1, s2, memo, m, n-1), editDistanceRecurse(s1, s2, memo, m-1, n),
			editDistanceRecurse(s1, s2, memo, m-1, n-1))
	}
	return memo[m][n]
}

// EditDistanceTabulation func solves the minimum number of operations required to convert s1 to s2 using
// insert, delete and replace, iteratively using dynamic programing
func EditDistanceTabulation(s1, s2 string, m, n int) int {
	memo := buildMemoTable(m+1, n+1)
	memo[0][0] = 0
	for i := 1; i <= n; i++ {
		memo[0][i] = i
	}
	for i := 1; i <= m; i++ {
		memo[i][0] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = 1 + min(memo[i-1][j], memo[i][j-1], memo[i-1][j-1])
			}
		}
	}
	return memo[m][n]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}
