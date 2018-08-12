func threeSum(nums []int) [][]int {
    out := make([][]int, 0)
    if len(nums) < 3 {
        return out
    }
    sort.Ints(nums)
    for i, _ := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j := i + 1
        k := len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum > 0 {
                k--
            } else if sum < 0 {
                j ++
            } else {
                ant := []int{nums[i], nums[j], nums[k]}
                out = append(out, ant)
                j++
                k--
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
                for j < k && nums[k] == nums[k+1] {
                    k--
                }
            }
        }
    }
    return out
}