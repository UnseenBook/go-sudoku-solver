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
	fmt.Println(localBoard)

	for localBoard.SetValuesBasedOnPossibilities() {
		localBoard.CalculatePossibilities()
		fmt.Println()
		fmt.Println(localBoard)
	}

	fmt.Println()
	fmt.Println("Finished puzzel:")
	fmt.Println(localBoard)
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
