package dp

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	want := 2
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	if got := uniquePathsWithObstacles(obstacleGrid); want != got {
		t.Errorf("uniquePathsWithObstacles(%v) = %d, want %d", obstacleGrid, got, want)
	}
	want = 1
	obstacleGrid = [][]int{{0}}
	if got := uniquePathsWithObstacles(obstacleGrid); want != got {
		t.Errorf("uniquePathsWithObstacles(%v) = %d, want %d", obstacleGrid, got, want)
	}
	want = 0
	obstacleGrid = [][]int{{1, 0}}
	if got := uniquePathsWithObstacles(obstacleGrid); want != got {
		t.Errorf("uniquePathsWithObstacles(%v) = %d, want %d", obstacleGrid, got, want)
	}
}
