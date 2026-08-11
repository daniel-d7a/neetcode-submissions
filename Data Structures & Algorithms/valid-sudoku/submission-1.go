func isValidSudoku(board [][]byte) bool {

	// check rows and cols
	for c, col := range board {
		rowNums := make(map[string]bool)
		colNums := make(map[string]bool)
		for r, row := range col {
			rowVal := string(row)
			colVal := string(board[r][c])
			if rowVal != "." && rowNums[rowVal] {
				return false
			}
			if colVal != "." && colNums[colVal] {
				return false
			}
			rowNums[rowVal] = true
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