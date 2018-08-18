func majorityElement(nums []int) int {
	sort.Ints(nums)
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}
	var v int
	var w int
	for i, x := range cnt {
		if x > v {
			v = x
			w = i
		}
	}
	return w
}