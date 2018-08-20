func judgeCircle(moves string) bool {
	opts := map[string]int{"U": 0, "D": 0, "L": 0, "R": 0}
	for _, x := range moves {
		opts[string(x)]++
	}
	if opts["U"] != opts["D"] || opts["L"] != opts["R"] {
		return false
	}
	return true
}