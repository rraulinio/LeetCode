func isPalindrome(x int) bool {
    y := x
    z := 0
    for y > 0 {
        z = z*10 + y%10
	    y /= 10
    }
    if x == z {
        return true
    }
    return false
}
