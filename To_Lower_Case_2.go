func toLowerCase(str string) string {
	var iis []string
	for _, x := range str {
		if x >= 65 && x <= 90 {
			x += 32 // what about some ascii?
		}
		iis = append(iis, string(x))
	}
	return strings.Join(iis, "")
}