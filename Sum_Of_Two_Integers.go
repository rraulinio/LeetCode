func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	carry := (a & b) << 1
	sum := a ^ b
	return getSum(sum, carry)
}