func findWords(words []string) []string {
	var out []string
	lits := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for _, x := range words {
		c := 0
		for _, y := range lits {
			if strings.ContainsAny(y, strings.ToLower(x)) {
				c++
			}
		}
		if c == 1 {
			out = append(out, x)
		}
	}
	return out
}