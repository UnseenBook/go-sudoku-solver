package board

// Board holds all the items of the Sudoku puzzel
type Board struct {
	items [9][9]boardItem
}

type boardItem struct {
	possibilities []int
	Value         int
}

func buildBoardFromInput(boardInput [][]int) {
	var localItems [9][9]boardItem
	for lineNum, line := range boardInput {
		for columnNum, item := range line {
			possi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			localItems[lineNum][columnNum] = boardItem{possi, item}
		}
	}
}
