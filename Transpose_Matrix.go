func transpose(A [][]int) [][]int {
	var B [][]int
	for i := 0; i < len(A[0]); i++ {
		var vls []int
		for j := 0; j < len(A); j++ {
			vls = append(vls, A[j][i])
		}
		B = append(B, vls)
	}
	return B
}