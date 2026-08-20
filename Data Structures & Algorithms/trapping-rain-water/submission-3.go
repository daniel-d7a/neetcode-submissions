func trap(height []int) int {

	total := 0

	l := 0
	r := len(height) - 1

	maxLeft := height[l]
	maxRight := height[r]

	for l < r {
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, height[l])
			total = total + maxLeft - height[l]
		} else {
			r--
			maxRight = max(maxRight, height[r])
			total = total + maxRight - height[r]
		}
	}


	return total
}