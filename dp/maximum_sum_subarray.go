package dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	memo[0] = nums[0]
	ans := -1 << 63
	for i := 1; i < len(nums); i++ {
		current := memo[i-1] + nums[i]
		memo[i] = max(current, nums[i])
		ans = max(memo[i], ans)
	}
	return max(memo[0], ans)
}
