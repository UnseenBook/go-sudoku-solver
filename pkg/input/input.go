package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func readLine() (string, error) {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if len(line) != 10 {
		return "", errors.New("input a line with 9 characters, no more, no less")
	}

	return line, nil
}

// GetSudokuBoard reads console input and build the board from it
func GetSudokuBoard() ([9][9]int, error) {
	var boardInput [9][9]int

	for count := 0; count < 9; count++ {
		line, err := readLine()
		if err != nil {
			return boardInput, err
		}

		for index, value := range line {
			if index == 9 {
				break
			}
			if value == ' ' {
				continue
			}
			if value < '0' || value > '9' {
				return boardInput, errors.New("only fill in numbers 0-9 and spaces")
			}
			boardInput[count][index] = int(value - '0')
		}
	}

	return boardInput, nil
}
