package dp

// CoinChangeNaive function calculates the number of ways the denominations can be used to reach the target value.
func CoinChangeNaive(array []int, sum int) int {
	return coinChangeNaive(array, len(array), sum)
}

func coinChangeNaive(array []int, length, sum int) int {
	if sum == 0 {
		return 1
	}
	if length == 0 {
		return 0
	}
	res := coinChangeNaive(array, length-1, sum)
	if array[length-1] <= sum {
		res += coinChangeNaive(array, length, sum-array[length-1])
	}
	return res
}

// CoinChange function calculates the number of ways the denominations can be used to reach the target value.
func CoinChange(coins []int, sum int) int {
	memo := buildMemoTable(sum+1, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		memo[0][i] = 1
	}
	for i := 1; i <= sum; i++ {
		memo[i][0] = 0
	}
	for i := 1; i <= sum; i++ {
		for j := 1; j <= len(coins); j++ {
			memo[i][j] = memo[i][j-1]
			if coins[j-1] <= i {
				memo[i][j] += memo[i-coins[j-1]][j]
			}
		}
	}
	return memo[sum][len(coins)]
}

func buildMemoTable(row, col int) [][]int {
	memo := make([][]int, row)
	for i := 0; i < row; i++ {
		memo[i] = make([]int, col)
	}
	return memo
}
