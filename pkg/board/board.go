package board

// Board holds all the items of the Sudoku puzzel
type Board struct {
	items [9][9]boardItem
}

type boardItem struct {
	possibilities []int
	Value         int
}

func newFilledBoardItem(value int) boardItem {
	return boardItem{possibilities: []int{value}, Value: value}
}

func newEmptyBoardItem() boardItem {
	return boardItem{possibilities: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
}

func buildBoardFromInput(boardInput [][]int) Board {
	var localItems [9][9]boardItem
	for lineNum, line := range boardInput {
		for collumnNum, item := range line {
			if item == 0 {
				localItems[lineNum][collumnNum] = newEmptyBoardItem()
				continue
			}
			localItems[lineNum][collumnNum] = newFilledBoardItem(item)
		}
	}

	return Board{items: localItems}
}

func (b *Board) setProperPossibilities() {
	for lineNum, line := range b.items {
		for collumnNum, item := range line {
			item.possibilities = getPossibilities(*b, lineNum, collumnNum)
		}
	}
}

func getPossibilities(b Board, lineNum int, collumnNum int) []int {
	possi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	return possi
}
