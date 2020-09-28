package array

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	return findThreeSumClosest(nums, target)
}

func findThreeSumClosest(nums []int, target int) int {
	diff := 1 << 32
	for i := 0; i < len(nums); i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if math.Abs(float64(target-sum)) < math.Abs(float64(diff)) {
				diff = target - sum
				if diff == 0 {
					return target
				}
			}
			if sum < target {
				low++
			} else {
				high--
			}
		}
	}
	return target - diff
}
