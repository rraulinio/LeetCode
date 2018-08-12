func intToRoman(num int) string {
	romans := make(map[int]string)
	romans[1] = "I"
	romans[5] = "V"
	romans[10] = "X"
	romans[50] = "L"
	romans[100] = "C"
	romans[500] = "D"
	romans[1000] = "M"
	romans[4] = "IV"
	romans[9] = "IX"
	romans[40] = "XL"
	romans[90] = "XC"
	romans[400] = "CD"
	romans[900] = "CM"
	out := ""
	x := num
	if v, ok := romans[x]; ok {
		return v
	}
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4}
	for _, val := range vals {
		if val > x {
			continue
		}
		for x/val > 0 {
			out += romans[val]
			x -= val
		}
	}
	if x > 0 {
		out += strings.Repeat(romans[1], x)
	}
	return out
}