func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	if strings.Compare(strings.ToLower(haystack), strings.ToLower(needle)) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}
		pos := i + len(needle)
		if strings.Compare(strings.ToLower(haystack[i:pos]), strings.ToLower(needle)) == 0 {
			return i
		}
	}
	return -1
}