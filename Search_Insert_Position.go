func searchInsert(nums []int, target int) int {
	out := len(nums)
	for i, _ := range nums {
		if nums[i] == target || nums[i] > target {
			out = i
			break
		}
	}
	return out
}