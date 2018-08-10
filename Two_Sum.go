func twoSum(nums []int, target int) []int {
    var out []int
    for i, v1 := range nums {
        for j, v2 := range nums {
            if (j > i) && (v1 + v2 == target) {
                out = append(out, i)
                out = append(out, j)
                break
            }
        }
    }
    return out
}
