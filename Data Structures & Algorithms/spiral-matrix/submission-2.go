func spiralOrder(matrix [][]int) []int {
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	output := []int{}

	for left <= right && top <= bottom{

		// to right
		for i := left; i <= right; i++ {
			output = append(output, matrix[top][i])
		}
		top++
		// to bottom
		for i := top; i <= bottom; i++ {
			output = append(output, matrix[i][right])
		}
		right--

		if top > bottom || left > right {
			break
		}

		// to left
		for i := right; i >= left; i-- {
			output = append(output, matrix[bottom][i])
		}
		bottom--
		// to top
		for i := bottom; i >= top; i-- {
			output = append(output, matrix[i][left])
		}
		left++
	}

	return output
}
