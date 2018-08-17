func numJewelsInStones(J string, S string) int {
	var out int
	for _, x := range S {
		exist := func(str string) bool {
			for _, x := range J {
				if string(x) == str {
					return true
				}
			}
			return false
		}(string(x))
		if exist {
			out++
		}
	}
	return out
}