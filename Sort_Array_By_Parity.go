func helper(A []int, B *[]int) {
	if len(A) == 0 {
		return
	}
	if A[0]%2 == 0 {
		*B = append([]int{A[0]}, *B...)
	} else {
		*B = append(*B, A[0])
	}
	helper(A[1:], B)
}
func sortArrayByParity(A []int) []int {
	var B []int
	helper(A, &B)
	return B
}