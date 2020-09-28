package array

import "sort"

type triplet struct {
	first  int
	second int
	third  int
}

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	return findTripletThatSumEqualsZero(nums)
}

func findTripletThatSumEqualsZero(nums []int) [][]int {
	allTriplets := make(map[triplet]bool)
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		previous := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			left := (target + nums[j]) * -1
			if _, ok := previous[left]; ok {
				triplets := triplet{first: nums[i], second: left, third: nums[j]}
				if _, has := allTriplets[triplets]; !has {
					allTriplets[triplets] = true
				}
			}
			previous[nums[j]] = j
		}
	}
	var ans [][]int
	for triplets := range allTriplets {
		ans = append(ans, []int{triplets.first, triplets.second, triplets.third})
	}
	return ans
}
