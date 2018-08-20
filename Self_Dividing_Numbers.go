func selfDividingNumbers(left int, right int) []int {
	var out []int
	for i := left; i <= right; i++ {
		str := fmt.Sprintf("%d", i)
		ok := func(i int, str string) bool {
			for _, x := range str {
				if string(x) == "0" {
					return false
				}
				var k int
				if _, err := fmt.Sscan(string(x), &k); err == nil && i%k > 0 {
					return false
				}
			}
			return true
		}(i, str)
		if ok {
			out = append(out, i)
		}
	}
	return out
}