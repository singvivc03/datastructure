package array

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	indexes := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		left := target - nums[i]
		if _, ok := indexes[left]; ok {
			return []int{indexes[left], i}
		}
		indexes[nums[i]] = i
	}
	return nil
}
