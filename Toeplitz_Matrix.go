func isToeplitzMatrix(matrix [][]int) bool {
	i, j := len(matrix)-1, 0
	for {
		x := matrix[i][j]
		a, b := i-1, j-1
		for a >= 0 && b >= 0 {
			y := matrix[a][b]
			if x != y {
				return false
			}
			a, b = a-1, b-1
		}
		a, b = i+1, j+1
		for a < len(matrix) && b < len(matrix[0]) {
			y := matrix[a][b]
			if x != y {
				return false
			}
			a, b = a+1, b+1
		}
		if i-1 >= 0 {
			i = i - 1
		} else if j+1 < len(matrix[0]) {
			j = j + 1
		} else {
			break
		}
	}
	return true
}