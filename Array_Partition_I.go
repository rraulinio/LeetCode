func arrayPairSum(nums []int) int {
    var number int
	sort.Ints(nums)
    for i := 0; i < len(nums); i += 2 {
        number += nums[i]
    }
    return number
}