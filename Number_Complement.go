// Could have used directly bits operations, but did not want to :)
func findComplement(num int) int {
	x := fmt.Sprintf("%b", num)
	var y string
	for _, v := range x {
		j, err := strconv.ParseInt(string(v), 2, 32)
		if err != nil {
			return 0
		}
		y += fmt.Sprintf("%d", ((j + 1) % 2))
	}
	n, err := strconv.ParseInt(y, 2, 32)
	if err != nil {
		return 0
	}
	return int(n)
}