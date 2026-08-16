func maxArea(heights []int) int {

	max := 0

	l := 0
	r := len(heights) - 1

	for l < r {

		min := 0
		if heights[l] < heights[r] {
			min = heights[l]
		} else {
			min = heights[r]
		}

		area := min * (r - l)

		if area > max {
			max = area
		}

		if heights[l] < heights[r]{
			l++
		} else {
			r--
		}

	}
	return max

}
