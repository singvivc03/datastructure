package array

import (
	"math"
)

func maxArea(height []int) int {
	low, high, area := 0, len(height)-1, 0
	for low <= high {
		heightToConsider := math.Min(float64(height[low]), float64(height[high]))
		area = int(math.Max(float64(area), heightToConsider*float64(high-low)))
		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return area
}
