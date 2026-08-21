func characterReplacement(s string, k int) int {

	charMap := make(map[string]int)

	l := 0
	r := 0

	answer := 0

	for r < len(s) {

		charMap[string(s[r])]++

		windowLength := r - l + 1
		maxFreq := getMaxFreq(charMap)
		result := windowLength - maxFreq

		if result <= k {
			answer = max(answer, windowLength)
		} else {
			for r - l + 1 - getMaxFreq(charMap) > k{
				charMap[string(s[l])]--
				l++
			}
		}
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