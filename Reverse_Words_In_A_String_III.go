func reverseWords(s string) string {
	v := strings.Split(s, " ")
	for i, x := range v {
		y := []rune(x)
		for a, b := 0, len(y)-1; a < b; a, b = a+1, b-1 {
			y[a], y[b] = y[b], y[a]
		}
		v[i] = string(y)
	}
	out := strings.Join(v, " ")
	return out
}