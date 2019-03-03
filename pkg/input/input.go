package input

import (
	"bufio"
	"fmt"
	"os"
)

// GetSudokuBoard reads console input and build the board from it
func GetSudokuBoard() ([9][9]int, error) {
	var boardInput [9][9]int

	allInput, err := readAllInput()
	if err != nil {
		return boardInput, err
	}
	for i := 0; i < 9; i++ {
		startInput := i * 9
		copy(boardInput[i][:], allInput[startInput:startInput+9])
	}

	return boardInput, nil
}

func readAllInput() ([]int, error) {
	in := bufio.NewReader(os.Stdin)
	var input []int
	for len(input) < 81 {
		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		input = mergeInputs(input, getNumbersFromLine(line))
	}

	return input, nil
}

func mergeInputs(left []int, right []int) []int {
	list := make([]int, len(left)+len(right))
	copy(list, left)
	copy(list[len(left):], right)

	return list
}

func getNumbersFromLine(line string) []int {
	var numbers []int
	for _, value := range line {
		if value == ' ' {
			numbers = append(numbers, 0)
		}
		if value >= '0' && value <= '9' {
			numbers = append(numbers, int(value-'0'))
		}
	}

	return numbers
}
