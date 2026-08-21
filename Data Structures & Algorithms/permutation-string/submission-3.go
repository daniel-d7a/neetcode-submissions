func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	count1 := make(map[rune]int)
	for _, c := range s1 {
		count1[c]++
	}

	l := 0
	r := len(s1)
	
	var isValid bool

	for r <= len(s2) {
		isValid = true
		count2 := make(map[rune]int)

		for _, c := range s2[l:r]{
			count2[c]++
		}

		for key, val := range count1 {
			if count2[key] != val {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}

		l++
		r++
	}

	return isValid
}
