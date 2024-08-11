package gamestate

func CheckWinner(board [3][3]string, player string) bool {
	// check the rows
	for i := 0; i < len(board); i++ {
		count := 0
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == player {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	// check the cols
	for j := 0; j < len(board[0]); j++ {
		count := 0
		for i := 0; i < len(board); i++ {
			if board[i][j] == player {
				count++
			}
		}

		if count == 3 {
			return true
		}
	}

	// check diagonals left to right
	count := 0
	for i := 0; i < len(board); i++ {
		if board[i][i] == player {
			count++
		}
	}

	if count == 3 {
		return true
	}

	// check diagonals right to left
	count = 0
	for i := 0; i < len(board); i++ {
		if board[i][2-i] == player {
			count++
		}
	}

	return count == 3
}

func CheckDraw(board [3][3]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}

	return true
}
