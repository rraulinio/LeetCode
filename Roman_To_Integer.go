func romanToInt(s string) int {
	romans := make(map[string]int)
	romans["I"] = 1
	romans["V"] = 5
	romans["X"] = 10
	romans["L"] = 50
	romans["C"] = 100
	romans["D"] = 500
	romans["M"] = 1000
	romans["IV"] = 4
	romans["IX"] = 9
	romans["XL"] = 40
	romans["XC"] = 90
	romans["CD"] = 400
	romans["CM"] = 900
	pos := 0
	out := 0
	for pos < len(s) {
		val := string(s[pos])
		if pos < len(s)-1 {
			val += string(s[pos+1])
		}
		if _, ok := romans[val]; ok {
			out += romans[val]
			pos += 2
		} else {
			val = string(s[pos])
			out += romans[val]
			pos += 1
		}
	}
	return out
}
