package board

import (
	"fmt"
	"strconv"
)

// Board holds all the items of the Sudoku puzzel
type Board [9][9]Item

// Item holds the data for a square in the puzzel
type Item struct {
	Possibilities sudokuList
	Value         int
}

type sudokuList []int

// SetValuesBasedOnPossibilities sets the Value of an item to the only possiblitie there is
func (b *Board) SetValuesBasedOnPossibilities() bool {
	anyUpdated := false

	for rowNum, line := range b {
		for collumnNum := range line {
			item := &b[rowNum][collumnNum]
			if item.Value == 0 && len(item.Possibilities) == 1 {
				anyUpdated = true
				item.Value = item.Possibilities[0]
			}
		}
	}

	return anyUpdated
}

// CalculateAndUpdatePossibilities calculates the possibilieties of the empty fields
func (b *Board) CalculateAndUpdatePossibilities() {
	for rowNum, line := range b {
		for collumnNum := range line {
			item := &b[rowNum][collumnNum]
			if item.Value == 0 {
				item.Possibilities = b.getPossibilitiesForSquare(collumnNum, rowNum)
			}
		}
	}
}

func (b *Board) getPossibilitiesForSquare(collumnNum int, rowNum int) []int {
	localPossibilities := subtractSudokuLists(completeSudokuList(), b.getBigSqaureValuesFromSquare(rowNum, collumnNum))
	localPossibilities = subtractSudokuLists(localPossibilities, b.getRowValuesFromSquare(rowNum, collumnNum))
	localPossibilities = subtractSudokuLists(localPossibilities, b.getCollumnValuesFromSquare(rowNum, collumnNum))

	return localPossibilities
}

func subtractSudokuLists(leftList sudokuList, rightList sudokuList) sudokuList {
	subtracted := make(sudokuList, 0, len(leftList))

	for _, possValue := range leftList {
		if !inSudokuList(possValue, rightList) {
			subtracted = append(subtracted, possValue)
		}
	}

	return subtracted
}

func addSudokuLists(leftList sudokuList, rightList sudokuList) sudokuList {
	localList := make(sudokuList, len(leftList), len(leftList)+len(rightList))
	copy(localList, leftList)
	for _, newVal := range rightList {
		if !inSudokuList(newVal, localList) {
			localList = append(localList, newVal)
		}
	}

	return localList
}

func inSudokuList(search int, haystack sudokuList) bool {
	for _, value := range haystack {
		if value == search {
			return true
		}
	}

	return false
}

func (b *Board) getBigSqaureValuesFromSquare(rowNum int, collumnNum int) sudokuList {
	var squareList sudokuList
	rowStart := rowNum / 3 * 3
	collumnStart := collumnNum / 3 * 3
	for rowIndex := rowStart; rowIndex < rowStart+3; rowIndex++ {
		for collumnIndex := collumnStart; collumnIndex < collumnStart+3; collumnIndex++ {
			if rowIndex == rowNum && collumnIndex == collumnNum {
				continue
			}
			if b[rowIndex][collumnIndex].Value != 0 {
				squareList = append(squareList, b[rowIndex][collumnIndex].Value)
			}
		}
	}

	return squareList
}

func (b *Board) getRowValuesFromSquare(rowNum int, collumnNum int) sudokuList {
	var line sudokuList
	for i := 0; i < 9; i++ {
		if b[rowNum][i].Value != 0 {
			if collumnNum == i {
				continue
			}
			line = append(line, b[rowNum][i].Value)
		}
	}

	return line
}

func (b *Board) getCollumnValuesFromSquare(rowNum int, collumnNum int) sudokuList {
	var collumn sudokuList
	for i := 0; i < 9; i++ {
		if b[i][collumnNum].Value != 0 {
			if rowNum == i {
				continue
			}
			collumn = append(collumn, b[i][collumnNum].Value)
		}
	}

	return collumn
}

func (b Board) String() string {
	var print string
	for rowNum, line := range b {
		for columnNum, item := range line {
			print += " "
			if item.Value == 0 {
				print += " "
			} else {
				print += strconv.Itoa(item.Value)
			}
			if columnNum == 2 || columnNum == 5 {
				print += " |"
			}
		}
		print += "\n"
		if rowNum == 2 || rowNum == 5 {
			print += "-------+-------+-------\n"
		}
	}

	return print
}

// BuildBoardFromInput creates a Board based on the input
func BuildBoardFromInput(boardInput [9][9]int) Board {
	var localItems [9][9]Item
	for rowNum, line := range boardInput {
		for collumnNum, item := range line {
			if item == 0 {
				localItems[rowNum][collumnNum] = newEmptyBoardItem()
				continue
			}
			localItems[rowNum][collumnNum] = newFilledBoardItem(item)
		}
	}

	return Board(localItems)
}

// IsSolved checks if the board is filled and correct
func (b *Board) IsSolved() bool {
	for _, line := range b {
		for _, item := range line {
			if item.Value == 0 {
				return false
			}

		}
	}
	return true
}

// ValidateBoard checks if the board has an error
func (b *Board) ValidateBoard() error {
	for rowNum, line := range b {
		for collumnNum := range line {
			return b.validateSquare(collumnNum, rowNum)
		}
	}

	return nil
}

func (b *Board) validateSquare(collumnNum int, rowNum int) error {
	if len(b.getPossibilitiesForSquare(collumnNum, rowNum)) == 0 {
		return fmt.Errorf("square has no possibilities. Coll: %v, Row: %v", collumnNum, rowNum)
	}

	return nil
}

func completeSudokuList() sudokuList {
	return sudokuList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func newFilledBoardItem(value int) Item {
	return Item{Possibilities: sudokuList{value}, Value: value}
}

func newEmptyBoardItem() Item {
	return Item{Possibilities: completeSudokuList()}
}

func getSudokuListIntersect(left sudokuList, right sudokuList) sudokuList {
	var intersect sudokuList
	for _, leftValue := range left {
		for _, rightValue := range right {
			if leftValue == rightValue {
				intersect = append(intersect, leftValue)
			}
		}
	}

	return intersect
}
