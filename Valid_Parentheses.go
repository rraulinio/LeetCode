func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	left := map[string]int{"{": 1, "[": 2, "(": 3}
	right := map[string]int{"}": 1, "]": 2, ")": 3}
	ol := []int{}
	for _, x := range s {
		v, ok := left[string(x)]
		if ok {
			ol = append([]int{v}, ol...)
		} else {
			if len(ol) == 0 {
				return false
			}
			var y int
			y, ol = ol[0], ol[1:]
			v = right[string(x)]
			if v != y {
				return false
			}
		}
	}
	if len(ol) > 0 {
		return false
	}
	return true
}