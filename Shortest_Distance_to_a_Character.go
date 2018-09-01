func shortestToChar(S string, C byte) []int {
	var out, pos []int
	for i, x := range S {
		if string(x) == string(C) {
			pos = append(pos, i)
		}
	}
	p := 0
	for i, x := range S {
		if p > 0 && p < len(pos) && math.Abs(float64(pos[p-1]-i)) < math.Abs(float64(pos[p]-i)) {
			out = append(out, int(math.Abs(float64(pos[p-1]-i))))
		} else {
			out = append(out, int(math.Abs(float64(pos[p]-i))))
		}
		if string(x) == string(C) && p < len(pos)-1 {
			p++
		}
	}
	return out
}