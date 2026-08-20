func lengthOfLongestSubstring(s string) int {

	l := 0
	r := 0

	charMap := make(map[byte]bool)

	maxWidth := 0

	for r < len(s) {
		for charMap[s[r]] {
			charMap[s[l]] = false
			l++
		}
		charMap[s[r]] = true
		maxWidth = max(maxWidth, r - l + 1)
		r++
	}

	return maxWidth










	// total := 0
	// counter := 0
	// charMap := make(map[rune]bool)

	// for _, val := range s {
	// 	if charMap[val] {
	// 		total = max(total, counter)
	// 		charMap = make(map[rune]bool)
	// 		counter = 0
	// 	} 
		
	// 	counter++
	// 	charMap[val] = true
	// }

	
	// total = max(total, counter)
	// return total

}
