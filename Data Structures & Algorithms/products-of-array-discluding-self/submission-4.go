func productExceptSelf(nums []int) []int {

	result := make([]int, len(nums))
	total := 1

	hasZero := false
	manyZero := false

	for _, val := range nums {
		if val == 0 && hasZero {
			manyZero = true
			total = 0
			break
		}else if val == 0 {
			hasZero = true
			continue
		}
		total = total * val
	}

	for i, val := range nums {
		if manyZero {
			result[i] = 0
		}else if hasZero {
			if val == 0{
				result[i] = total
			}else{
				result [i] = 0
			}
		}else{
			result[i] = total / val
		}
	}
	return result

}
