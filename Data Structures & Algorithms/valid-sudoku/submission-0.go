func isValidSudoku(board [][]byte) bool {

	// check rows
	for _, col := range board {
		rowNums := make(map[string]bool)

		for _, row := range col {
			rowVal := string(row)
			if rowVal != "." && rowNums[rowVal] {
				return false
			}
			rowNums[rowVal] = true
		}
	}

	// // check cols
	for c, col := range board {
		colNums := make(map[string]bool)

		for r := range col {

			colVal := string(board[r][c])

			if colVal != "." && colNums[colVal] {
				return false
			}
			colNums[colVal] = true
		}
	}

	// check grid squares
	gridLen := int(math.Sqrt(float64(len(board))))

	for j := 0; j < gridLen; j++ {
		for i := 0; i < gridLen; i++ {

			gridNums := make(map[string]bool)

			for y := 0; y < gridLen; y++ {
				for x := 0; x < gridLen; x++ {
					xPos := x + gridLen * i
					yPos := y + gridLen * j

					gridVal := string(board[xPos][yPos])

					if gridVal != "." && gridNums[gridVal] {
						return false
					}
					gridNums[gridVal] = true

				}
			}
		}
	}

	return true

}