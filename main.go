package main

import (
	"fmt"
	"math/rand/v2"

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
	currentPlayer := rand.IntN(2)
	if currentPlayer == 0 {
		msg := fmt.Sprintln("Computer will start the game")
		color.Green(msg)
	} else {
		msg := fmt.Sprintln("Player will start the game")
		color.Green(msg)
	}

	for {
		printBoard(board)
		var row, col int

		switch players[currentPlayer] {
		case "X":
			fmt.Println("computer's turn")
			row, col = gamelogic.GetComputerMove(board)
		default:
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
