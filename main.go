package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printBoard(board [3][3]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println(" " + join(board[i], "  |  "))
		if i < 2 {
			fmt.Println("----------------")
		}
	}
}

func join(elements [3]string, separator string) string {
	result := ""
	for index, element := range elements {
		result += element
		if index < len(elements)-1 {
			result += separator
		}
	}
	return result
}

func getMove(player string) int {
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

	return move
}

func divmod3(num int) (int, int) {
	return num / 3, num % 3
}

func checkWinner(board [3][3]string, player string) bool {
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

func checkDraw(board [3][3]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}

	return true
}

func main() {
	board := [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	players := [2]string{"X", "O"}
	currentPlayer := 0

	for {
		printBoard(board)
		move := getMove(players[currentPlayer])
		row, col := divmod3(move - 1)

		if board[row][col] == " " {
			board[row][col] = players[currentPlayer]
			if checkWinner(board, players[currentPlayer]) {
				msg := fmt.Sprintf("Player %s is the Winner!!", players[currentPlayer])
				color.Green(msg)
				printBoard(board)
				break
			} else if checkDraw(board) {
				msg := fmt.Sprintln("This game is a draw")
				color.Red(msg)
				printBoard(board)
				break
			}
			currentPlayer = (currentPlayer + 1) % 2
		} else {
			msg := fmt.Sprintln("spot already taken !!")
			color.Red(msg)
		}
	}
}
