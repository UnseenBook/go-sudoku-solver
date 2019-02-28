package main

import (
	"fmt"

	"github.com/UnseenBook/go-sudoku-solver/pkg/board"
	"github.com/UnseenBook/go-sudoku-solver/pkg/input"
)

func main() {
	inputBoard, err := input.GetSudokuBoard()

	if err != nil {
		fmt.Println(err)
	}

	localBoard := board.BuildBoardFromInput(inputBoard)

	localBoard.CalculatePossibilities()

	fmt.Println()
	printBoard(localBoard)

	for localBoard.SetValuesBasedOnPossibilities() {
		localBoard.CalculatePossibilities()
		fmt.Println()
		printBoard(localBoard)
	}

	fmt.Println()
	fmt.Println("Finished puzzel:")
	printBoard(localBoard)
}

func printBoardInput(board [9][9]int) {
	for rowNum, line := range board {
		for columnNum, item := range line {
			fmt.Print(" ")
			if item == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(item)
			}
			if columnNum == 2 || columnNum == 5 {
				fmt.Print(" |")
			}
		}
		fmt.Println()
		if rowNum == 2 || rowNum == 5 {
			fmt.Println(" - - - | - - - | - - -")
		}
	}
}

func printBoard(board board.Board) {
	for rowNum, row := range board {
		for columnNum, item := range row {
			fmt.Print(" ")
			if item.Value == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(item.Value)
			}
			// fmt.Println(item.Possibilities)
			if columnNum == 2 || columnNum == 5 {
				fmt.Print(" |")
			}
		}
		fmt.Println()
		if rowNum == 2 || rowNum == 5 {
			fmt.Println(" - - - | - - - | - - -")
		}
	}
}
