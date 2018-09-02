func isMonotonic(A []int) bool {
	ok := map[string]bool{"asc": true, "desc": true}
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			ok["asc"] = false
		}
	}
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			ok["desc"] = false
		}
	}
	if ok["asc"] || ok["desc"] {
		return true
	}
	return false
}