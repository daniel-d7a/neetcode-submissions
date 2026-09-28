func spiralOrder(matrix [][]int) []int {
	top := 0
	bottom := len(matrix) - 1
	left := -1
	right := len(matrix[0]) - 1

	output := []int{}

	n := len(matrix) * len(matrix[0])
	for j := 0; j < n; {

		// to right
		for i := left + 1; i <= right; i++ {
			output = append(output, matrix[top][i])
			j++
		}

		if j == n {
			break
		}

		// to bottom
		for i := top + 1; i <= bottom; i++ {
			output = append(output, matrix[i][right])
			j++
		}

		if j == n {
			break
		}
		left++
		// to left
		for i := right - 1; i >= left; i-- {
			output = append(output, matrix[bottom][i])
			j++
		}
		if j == n {
			break
		}

		// to top
		for i := bottom - 1; i > top; i-- {
			output = append(output, matrix[i][left])
			j++
		}
		if j == n {
			break
		}

		right--
		top++
		bottom--
	}

	return output
}
