package board

// Board holds all the items of the Sudoku puzzel
type Board [9][9]Item

// Item holds the data for a square in the puzzel
type Item struct {
	Possibilities possibilities
	Value         int
}

type sudokuList []int

type possibilities sudokuList

// SetValuesBasedOnPossibilities sets the Value of an item to the only possiblitie there is
func (b *Board) SetValuesBasedOnPossibilities() bool {
	anyUpdated := false

	for lineNum, line := range b {
		for collumnNum := range line {
			item := &b[lineNum][collumnNum]
			if item.Value == 0 && len(item.Possibilities) == 1 {
				anyUpdated = true
				item.Value = item.Possibilities[0]
			}
		}
	}

	return anyUpdated
}

// CalculatePossibilities calculates the possibilieties of the empty fields
func (b *Board) CalculatePossibilities() {
	var localPossibilities possibilities
	for lineNum, line := range b {
		for collumnNum := range line {
			item := &b[lineNum][collumnNum]
			if item.Value == 0 {
				localPossibilities = item.Possibilities.subtract(b.getBigSqaureValues(lineNum, collumnNum))
				localPossibilities = localPossibilities.subtract(b.getRowValues(lineNum))
				localPossibilities = localPossibilities.subtract(b.getCollumnValues(collumnNum))

				item.Possibilities = localPossibilities
			}
		}
	}
}

// BuildBoardFromInput creates a Board based on the input
func BuildBoardFromInput(boardInput [9][9]int) Board {
	var localItems [9][9]Item
	for lineNum, line := range boardInput {
		for collumnNum, item := range line {
			if item == 0 {
				localItems[lineNum][collumnNum] = newEmptyBoardItem()
				continue
			}
			localItems[lineNum][collumnNum] = newFilledBoardItem(item)
		}
	}

	return Board(localItems)
}

func completePossibilities() possibilities {
	return possibilities([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func (pRec possibilities) subtract(list sudokuList) possibilities {
	var subtracted possibilities
	found := false

	for _, possValue := range pRec {
		for _, subValue := range list {
			if possValue == subValue {
				found = true
				break
			}
		}
		if !found {
			subtracted = append(subtracted, possValue)
		}
		found = false
	}

	return subtracted
}

func newFilledBoardItem(value int) Item {
	return Item{Possibilities: possibilities{value}, Value: value}
}

func newEmptyBoardItem() Item {
	return Item{Possibilities: completePossibilities()}
}

func (b *Board) getBigSqaureValues(rowNum int, collumnNum int) sudokuList {
	var squareList sudokuList
	rowStart := rowNum / 3 * 3
	collumnStart := collumnNum / 3 * 3
	for rowIndex := rowStart; rowIndex < rowStart+3; rowIndex++ {
		for collumnIndex := collumnStart; collumnIndex < collumnStart+3; collumnIndex++ {
			if b[rowIndex][collumnIndex].Value != 0 {
				squareList = append(squareList, b[rowIndex][collumnIndex].Value)
			}
		}
	}

	return squareList
}

func (b *Board) getRowValues(rowNum int) sudokuList {
	var line sudokuList
	for i := 0; i < 9; i++ {
		if b[rowNum][i].Value != 0 {
			line = append(line, b[rowNum][i].Value)
		}
	}

	return line
}

func (b *Board) getCollumnValues(collumnNum int) sudokuList {
	var collumn sudokuList
	for i := 0; i < 9; i++ {
		if b[i][collumnNum].Value != 0 {
			collumn = append(collumn, b[i][collumnNum].Value)
		}
	}

	return collumn
}
