func hammingDistance(x int, y int) int {
	var cc int
    var i uint
	for i = 0; i < 32; i++ {
		if (uint(x^y) >> i) & 1 > 0 {
			cc++
		}
	}
	return cc
}