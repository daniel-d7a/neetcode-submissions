func isPalindrome(s string) bool {

	left := 0
	right := len(s) - 1

	for true {
		for left < len(s) && !isAlphanumeric(rune(s[left])) {
			left = left + 1
		}
		for right >= 0 && !isAlphanumeric(rune(s[right])) {
			right = right - 1
		}
		if right < 0 || left >= len(s) {
			break
		}

		leftStr := strings.ToLower(string(s[left]))
		rightStr := strings.ToLower(string(s[right]))

		if leftStr != rightStr {
			return false
		}
		left = left + 1
		right = right - 1
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}