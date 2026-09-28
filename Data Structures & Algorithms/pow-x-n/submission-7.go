func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	result := calcPow(x, int(math.Abs(float64(n))))
	if n < 0 {
		result = 1 / result
	}
	return result
}

func calcPow(x float64, n int) float64 {
	if n <= 1 {
		return x
	} else {
		res := calcPow(x, n/2)
		var result float64
		if n%2 == 0 {
			result = res * res
		} else {
			result = res * res * x
		}

		return result
	}
}