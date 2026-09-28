func plusOne(digits []int) []int {
    
	i := len(digits) - 1
	digits[i]++

	for i >= 0 && digits[i] > 9{
		digits[i] = 0

		if i == 0 {
			digits = append([]int{1}, digits...)
		} else {
			digits[i - 1]++
		}
		
		i--
	}

	// for i := len(digits) - 1; i >= 0; i-- {
	// 	if digits[i] < 9 {
	// 		digits[i]++
	// 	} else {
	// 		digits[i] = 0

	// 		if i == 0 {
	// 			digits = append([]int{0}, digits...)
	// 		} else {
	// 			digits[i - 1]++
	// 		}
	// 	}
	// } 

	return digits
}
