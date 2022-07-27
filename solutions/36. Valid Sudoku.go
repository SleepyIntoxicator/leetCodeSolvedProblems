package solutions

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowH, colH := make(map[byte]struct{}, 0), make(map[byte]struct{}, 0)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := colH[board[i][j]]; ok {
					return false
				} else {
					colH[board[i][j]] = struct{}{}
				}
			}
			if board[j][i] != '.' {
				if _, ok := rowH[board[j][i]]; ok {
					return false
				} else {
					rowH[board[j][i]] = struct{}{}
				}
			}

		}

		fromR, fromC := (i/3)*3, (i%3)*3
		squareDupls := make(map[byte]struct{})

		for r := fromR; r < fromR+3; r++ {
			for c := fromC; c < fromC+3; c++ {
				if board[r][c] == '.' {
					continue
				}
				if _, ok := squareDupls[board[r][c]]; ok {
					return false
				}
				squareDupls[board[r][c]] = struct{}{}
			}
		}
	}

	return true
}

// i - rows, vertical | j - cols, horizontal
