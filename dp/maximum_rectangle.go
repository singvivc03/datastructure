package dp

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	hist := make([]int, len(matrix[0]))
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				hist[j] = 0
			} else {
				hist[j] += int(matrix[i][j])
			}
		}
		area := maxArea(hist)
		if area > max {
			max = area
		}
	}
	return max
}

func maxArea(hist []int) int {
	indexStack := make([]int, len(hist))
	index, area := 1, -1<<63
	indexStack[0] = 0
	for i := 1; i < len(hist); {
		if index == 0 || hist[i] >= hist[indexStack[index-1]] {
			indexStack[index], index = i, index+1
			i = i + 1
		} else {
			top := indexStack[index-1]
			index = index - 1
			area = max(area, findArea(index, hist, indexStack, i, top))
		}
	}
	i := len(hist)
	for index != 0 {
		top := indexStack[index-1]
		index = index - 1
		area = max(area, findArea(index, hist, indexStack, i, top))
	}
	return area
}

func findArea(index int, hist []int, indexStack []int, i int, top int) int {
	var area int
	if index == 0 {
		area = hist[top] * i
	} else {
		area = hist[top] * (i - indexStack[index-1] - 1)
	}
	return area
}
