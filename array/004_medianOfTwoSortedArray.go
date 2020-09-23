package array

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	x, y, low, high := len(nums1), len(nums2), 0, len(nums1)
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := ((x + y + 1) / 2) - partitionX

		maxLeftX, minRightX := getMaxValue(partitionX, nums1), getMinValue(partitionX, x, nums1)
		maxLeftY, minRightY := getMaxValue(partitionY, nums2), getMinValue(partitionY, y, nums2)

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			median := math.Max(float64(maxLeftX), float64(maxLeftY))
			if (x+y)%2 == 0 {
				second := math.Min(float64(minRightX), float64(minRightY))
				median = (median + second) / 2
			}
			return median
		}
		if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}

func getMaxValue(partition int, nums []int) int {
	if partition == 0 {
		return -1 << 32
	}
	return nums[partition-1]
}

func getMinValue(partition int, max int, nums []int) int {
	if partition == max {
		return 1 << 32
	}
	return nums[partition]
}
