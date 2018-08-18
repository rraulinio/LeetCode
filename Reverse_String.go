func reverseString(s string) string {
    var out string
    for i := len(s) - 1; i >= 0; i-- {
        out += string(s[i])
    }
    return out
}