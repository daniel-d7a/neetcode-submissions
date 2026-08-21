func characterReplacement(s string, k int) int {

	charMap := make(map[string]int)

	l := 0
	r := 0

	answer := 0

	for r < len(s) {

		charMap[string(s[r])]++
		for r - l + 1 - getMaxFreq(charMap) > k{
			charMap[string(s[l])]--
			l++
		}

		answer = max(answer, r - l + 1)
		r++

	}

	return answer
}

func getMaxFreq(charMap map[string]int) int {
	maxFreq := 0
	for _, val := range charMap {
		maxFreq = max(maxFreq, val)
	}
	return maxFreq
}