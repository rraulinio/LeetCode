func fizzBuzz(n int) []string {
	var out []string
	for i := 1; i <= n; i++ {
		var v string
		if i % 3 == 0 && i % 5 == 0 {
			v = "FizzBuzz"
		} else if i % 3 == 0 {
			v = "Fizz"
		} else if i % 5 == 0 {
			v = "Buzz"
		} else {
			v = fmt.Sprintf("%d", i)
		}
		out = append(out, v)
	}
	return out
}