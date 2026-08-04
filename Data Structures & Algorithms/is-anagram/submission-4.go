func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	charSet := [26]int{}

	for i := range s {
		sSetIdx := int(s[i]) - 97
		tSetIdx := int(t[i]) - 97

		charSet[sSetIdx]++
		charSet[tSetIdx]--
	}

	for _, val := range charSet {
		if val != 0 {
			return false
		}
	}

	return true
}
