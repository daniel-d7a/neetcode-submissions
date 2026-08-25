func minWindow(s string, t string) string {
    
	if len(t) > len(s) {
		return ""
	}

	tMap := make(map[byte]int)

	for i := range t {
		tMap[t[i]]++
	}

	windowRight, windowLeft := -1, -1

	windowMap := make(map[byte]int)
	for l, r := 0, 0; r < len(s); r++ {
		windowMap[s[r]]++

		for checkValid(windowMap, tMap) {

			if windowLeft == -1 || r - l < windowRight - windowLeft {
				windowLeft, windowRight = l, r
			}

			windowMap[s[l]] = max(windowMap[s[l]] - 1, 0)
			l++
		}
	} 

	if windowLeft == -1 {
		return ""
	}
	return s[windowLeft : windowRight + 1]
}

func checkValid(window, sub map[byte]int) bool {
	for key, val := range sub {
		if window[key] < val {
			return false
		}
	}
	return true
}
