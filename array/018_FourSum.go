package array

import (
	"sort"
)

type quadruplets struct {
	first, second, third, fourth int
}

func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	return findFourSum(nums, target)
}

func findFourSum(nums []int, target int) [][]int {
	allQuadruplets := make(map[quadruplets]bool)
	for j := 0; j < len(nums); j++ {
		first := nums[j]
		for i := j + 1; i < len(nums); i++ {
			left := target - (first + nums[i])
			low, high := i+1, len(nums)-1
			for low < high {
				sum := nums[low] + nums[high]
				if sum == left {
					quadruplets := quadruplets{first: first, second: nums[i], third: nums[low], fourth: nums[high]}
					allQuadruplets[quadruplets] = true
					low++
				} else if sum < left {
					low++
				} else {
					high--
				}
			}
		}
	}
	var ans [][]int
	for key := range allQuadruplets {
		ans = append(ans, []int{key.first, key.second, key.third, key.fourth})
	}
	return ans
}
