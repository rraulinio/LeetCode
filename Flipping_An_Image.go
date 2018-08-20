func flipAndInvertImage(A [][]int) [][]int {
	for k, s := range A {
		for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
			if i == j {
				s[i] = (s[i]+1)%2
				break
			}
			s[i], s[j] = s[j], s[i]
			s[i] = (s[i]+1)%2
			s[j] = (s[j]+1)%2
		}
		A[k] = s
	}
	return A
}