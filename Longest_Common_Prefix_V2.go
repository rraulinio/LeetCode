func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    out := ""
    word := strs[0]
    ln := len(word)
    for _, x := range strs {
		if len(x) < ln {
            word = x
            ln = len(x)
		}
	}
    for i := 0; i < ln; i++ {
        stop := false
        for _, x := range strs {
            if x[i] != word[i] {
                stop = true
                break
            }
        }
        if stop {
            break
        }
        out += string(word[i])
    }
    return out
}