func setZeroes(matrix [][]int) {
	top_row := false

	for i := range matrix {
		for j := range matrix[i] {
			if matrix [i][j] == 0 {
				if i == 0 {
					top_row = true
				} else {
					matrix[i][0] = 0
				}
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {

			if (i == 0 && top_row) || (i != 0 && (matrix [0][j] == 0 || matrix[i][0] == 0)) {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := range matrix {
			matrix[j][0] = 0
		}
	}
	if top_row {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
}
