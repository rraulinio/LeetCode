func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    cp := strs
	out := ""
	si := 0
	word := ""
	for i, x := range cp {
		if len(x) <= len(cp[si]) {
			word = x
			si = i
		}
	}
	cp = append(cp[:si], cp[si+1:]...)
	for i := 0; i < len(word); i++ {
        if i > 0 && len(out) == 0 {
            break
        }
		ok := true
		for _, x := range cp {
			if x[i] != word[i] {
				ok = false
			}
		}
		if ok {
			out += string(word[i])
		}
	}
	return out
}
