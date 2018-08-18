func singleNumber(nums []int) int {
	var out int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == len(nums) - 1 {
			out = nums[i]
			break
		}
		if nums[i] == nums[i+1] {
			i++
			continue
		}
		out = nums[i]
	}
	return out
}