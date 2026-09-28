func spiralOrder(matrix [][]int) []int {
	top := 0
	bottom := len(matrix)
	left := 0
	right := len(matrix[0])

	output := []int{}

	for left < right && top < bottom{

		// to right
		for i := left; i < right; i++ {
			output = append(output, matrix[top][i])
		}
		top++
		// to bottom
		for i := top; i < bottom; i++ {
			output = append(output, matrix[i][right - 1])
		}
		right--

		if top >= bottom || left >= right {
			break
		}

		// to left
		for i := right - 1; i >= left; i-- {
			output = append(output, matrix[bottom - 1][i])
		}
		bottom--
		// to top
		for i := bottom - 1; i >= top; i-- {
			output = append(output, matrix[i][left])
		}
		left++
	}

	return output
}
