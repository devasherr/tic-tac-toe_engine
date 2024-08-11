package main

import (
	"fmt"

	gamelogic "github.com/devasherr/tic-tac-toe/gamelogic"
	gamestate "github.com/devasherr/tic-tac-toe/gamestate"
	"github.com/fatih/color"
)

func printBoard(board [3][3]string) {
	fmt.Println("")
	for i := 0; i < len(board); i++ {
		fmt.Println(" " + join(board[i], "  |  "))
		if i < 2 {
			fmt.Println("----------------")
		}
	}
	fmt.Println("")
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
		row, col := 0, 0

		if players[currentPlayer] == "X" {
			// computers turn
			fmt.Println("thinking...")
			row, col = gamelogic.GetComputerMove(board)
		} else {
			row, col = gamelogic.GetPlayerMove(players[currentPlayer])
		}

		if board[row][col] == " " {
			board[row][col] = players[currentPlayer]
			if gamestate.CheckWinner(board, players[currentPlayer]) {
				msg := fmt.Sprintf("Player %s is the Winner!!", players[currentPlayer])
				color.Green(msg)
				printBoard(board)
				break
			} else if gamestate.CheckDraw(board) {
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
