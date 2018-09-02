func calPoints(ops []string) int {
	var totals []int
	for _, sign := range ops {
		switch sign {
		case "+":
			v := totals[len(totals)-1] + totals[len(totals)-2]
			totals = append(totals, v)
		case "D":
			v := 2 * totals[len(totals)-1]
			totals = append(totals, v)
		case "C":
			totals = totals[0 : len(totals)-1]
		default:
			x, err := strconv.ParseInt(sign, 10, 64)
			if err != nil {
				log.Printf("Couldn't parse %s", sign)
				continue
			}
			totals = append(totals, int(x))
		}
	}
	var total int
	for i := range totals {
		total += totals[i]
	}
	return total
}