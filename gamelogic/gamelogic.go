package gamelogic

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/devasherr/tic-tac-toe/gamestate"
)

func GetPlayerMove(player string) (int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Player %s: input a number from 1 - 9:\n", player)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		move, err := strconv.Atoi(input)
		if err != nil || move < 1 || move > 9 {
			fmt.Println("Number must be from 1 - 9")
			continue
		}

		row, col := (move-1)/3, (move-1)%3
		return row, col
	}

}

func GetComputerMove(board [3][3]string) (int, int) {
	best_score := math.Inf(-1)
	row, col := 0, 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == " " {
				board[i][j] = "X"
				score := minimax(board, false)
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

func minimax(board [3][3]string, isMaximizing bool) float64 {
	if gamestate.CheckWinner(board, "X") {
		return 10
	}
	if gamestate.CheckWinner(board, "O") {
		return -10
	}
	if gamestate.CheckDraw(board) {
		return 0
	}

	if isMaximizing {
		bestScore := -math.Inf(1)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == " " {
					board[i][j] = "X"
					score := minimax(board, false)
					board[i][j] = " "
					bestScore = math.Max(bestScore, score)
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] == " " {
					board[i][j] = "O"
					score := minimax(board, true)
					board[i][j] = " "
					bestScore = math.Min(bestScore, score)
				}
			}
		}
		return bestScore
	}
}
