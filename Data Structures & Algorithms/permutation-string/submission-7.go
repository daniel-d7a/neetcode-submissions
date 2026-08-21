func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	count1 := make(map[rune]int)
	for _, c := range s1 {
		count1[c]++
	}

	var isValid bool

	for l, r := 0, len(s1); r <= len(s2); l, r = l+1, r+1 {
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
	}

	return false
}
