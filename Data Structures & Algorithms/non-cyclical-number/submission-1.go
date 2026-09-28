func isHappy(n int) bool {
    hash := make(map[int]bool)

	for n != 1 {
		n = getNumberSquaresSum(n)
		if hash[n] {
			return false
		}
		hash[n] = true
	}
	return true
}

func getNumberSquaresSum(n int) int {
	sum := 0
	for n != 0 {
		d := (n % 10)
		sum = sum + (d * d)
		n = n / 10
	}
	return sum
}