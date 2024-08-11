package gamelogic

import (
	"fmt"
	"math"

	"github.com/devasherr/tic-tac-toe/gamestate"
)

func GetPlayerMove(player string) (int, int) {
	var move int

	for {
		fmt.Printf("player %s: input a number from 1 - 9:\n", player)
		fmt.Scanf("%d", &move)
		if move <= 0 || move > 9 {
			fmt.Println("number must be from 1 - 9")
		} else {
			break
		}

	}

	return (move - 1) / 3, (move - 1) % 3
}

func GetComputerMove(board [3][3]string) (int, int) {
	best_score := math.Inf(-1)
	row, col := 0, 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == " " {
				board[i][j] = "X"
				score := minimax(board, 0, false)
				board[i][j] = " "
				if score > best_score {
					best_score = score
					row, col = i, j
				}
			}
		}
	}

	return row, col
}

func minimax(board [3][3]string, depth float64, is_maximizing bool) float64 {
	if gamestate.CheckWinner(board, "X") {
		return 10 - depth
	}
	if gamestate.CheckWinner(board, "O") {
		return depth - 10
	}
	if gamestate.CheckDraw(board) {
		return 0
	}

	if is_maximizing {
		best_score := math.Inf(-1)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == " " {
					board[i][j] = "X"
					score := minimax(board, depth+1, false)
					board[i][j] = " "
					best_score = max(best_score, score)
				}
			}
		}
		return best_score
	} else {
		best_score := math.Inf(-1)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == " " {
					board[i][j] = "X"
					score := minimax(board, depth+1, false)
					board[i][j] = "	 "
					best_score = max(best_score, score)
				}
			}
		}
		return best_score
	}
}
