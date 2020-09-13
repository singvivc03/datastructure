package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo := memoTable(len(obstacleGrid), len(obstacleGrid[0]), obstacleGrid)
	for rowIndex := 1; rowIndex < len(obstacleGrid); rowIndex++ {
		for colIndex := 1; colIndex < len(obstacleGrid[0]); colIndex++ {
			if obstacleGrid[rowIndex][colIndex] == 1 {
				continue
			}
			memo[rowIndex][colIndex] = memo[rowIndex-1][colIndex] + memo[rowIndex][colIndex-1]
		}
	}
	return memo[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func memoTable(rows int, cols int, obstacleGrid [][]int) [][]int {
	memo := make([][]int, rows)
	for rowIndex := range memo {
		memo[rowIndex] = make([]int, cols)
	}
	for i := 0; i < cols; i++ {
		if obstacleGrid[0][i] != 1 {
			if i == 0 {
				memo[0][i] = 1
			} else {
				memo[0][i] = memo[0][i-1]
			}
		}
	}
	for i := 1; i < rows; i++ {
		if obstacleGrid[i][0] != 1 {
			memo[i][0] = memo[i-1][0]
		}
	}
	return memo
}
