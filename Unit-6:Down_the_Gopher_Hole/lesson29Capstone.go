/*
	Things to implement
		1) Implement the following Rules
			a) number not in current row
			b) number not in current column
			c) number not in current 3x3 subregion
		2) Implement a method to set a digit at specific location
		3) Implement a method to remove a digit at specific location (set location to 0)
		4) Keep the program from changing fixed digits (non-zero digits form "mainSudoku")
*/
package main

import (
	"errors"
	"fmt"
)

type sudokuMap [row][columns]int8

const row = 9
const columns = 9

//Won't let me declare this as a const.
var mainSudoku = sudokuMap{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func (m sudokuMap) newSudoku() sudokuMap {
	return mainSudoku
}
func (m sudokuMap) checkRows(cRow, number int8) bool {
	for i := 0; i < columns; i++ {
		if m[cRow][i] == number {
			return true
		}
	}
	return false
}
func (m sudokuMap) checkColumns(cCol, number int8) bool {
	for i := 0; i < row; i++ {
		if m[i][cCol] == number {
			return true
		}
	}
	return false
}
func (m sudokuMap) findSubRegion(cRow, cCol int8) (subRegion int8, err error) {
	var (
		rowRegion string //Split up into 3 regions R1(0,2) R2(3,5) R3(6,8)
		colRegion string //Split up into 3 regions C1(0,2) C2(3,5) C3(6,8)
		//subRegion int8 //0 = (R1,C1), 1 =(R1,C2), 2 = (R1,C3)...8 = (R3,C3)
		//Will be used as an index for checkSubRegion
	)
	//I seperated this into 3 parts to make reading easier

	//Find Row Region (rowRegion)
	switch cRow {
	case 0, 1, 2:
		rowRegion = "R1"
	case 3, 4, 5:
		rowRegion = "R2"
	case 6, 7, 8:
		rowRegion = "R3"
	}
	//Find Column Region (colRegion)
	switch cCol {
	case 0, 1, 2:
		colRegion = "C1"
	case 3, 4, 5:
		rowRegion = "C2"
	case 6, 7, 8:
		rowRegion = "C3"
	}
	//Find SubRegion, setup this way for readability
	switch /*rowRegion*/ {
	case rowRegion == "R1":
		switch colRegion {
		case "C1":
			subRegion = 0 //(R1,C1)
		case "C2":
			subRegion = 1 //(R1,C2)
		case "C3":
			subRegion = 2 //(R1,C3)

		} //Regions 0-2
	case rowRegion == "R2":
		switch colRegion {
		case "C1":
			subRegion = 3 //(R1,C1)
		case "C2":
			subRegion = 4 //(R1,C2)
		case "C3":
			subRegion = 5 //(R1,C3)

		} //Regions 3-5
	case rowRegion == "R3":
		switch colRegion {
		case "C1":
			subRegion = 6 //(R1,C1)
		case "C2":
			subRegion = 7 //(R1,C2)
		case "C3":
			subRegion = 8 //(R1,C3)
		} //Regions 6-8
	default:
		err = errors.New("SubRegion Not Found")
		fmt.Println("Error Encountered: ", err)
	}
	return subRegion, err
} //err returns either nil or
// "Subregion Not Found
func (m sudokuMap) checkSubRegions(cRow, cCol, number int8) (check bool, err error) {
	subRegion, err := m.findSubRegion(cRow, cCol)
	if err != nil {
		fmt.Println("Error Encountered When Calling findSubregion()")
	}
	//This is the best that I could think. subRegions sorted by row then column and then are search
	switch subRegion {
	case 0, 1, 2: //Top Regions
		switch subRegion {
		case 0: //Left
			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 1: //Center
			for r := 0; r < 3; r++ {
				for c := 3; c < 6; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 2: //Right
			for r := 0; r < 3; r++ {
				for c := 6; c < 9; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		}
	case 3, 4, 5: //Middle Regions
		switch subRegion {
		case 3: //Left
			for r := 3; r < 6; r++ {
				for c := 0; c < 3; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 4: //Center
			for r := 3; r < 6; r++ {
				for c := 3; c < 6; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 5: //Right
			for r := 3; r < 6; r++ {
				for c := 6; c < 9; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		}
	case 6, 7, 8: //Bottom Regions
		switch subRegion {
		case 6: //Left
			for r := 6; r < 9; r++ {
				for c := 0; c < 3; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 7: //Center
			for r := 6; r < 9; r++ {
				for c := 3; c < 6; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		case 8: //Right
			for r := 6; r < 9; r++ {
				for c := 6; c < 9; c++ {
					if number == m[r][c] {
						check = true
					}
				}
			}
		default:
			err = errors.New("Subregion: not found")
		}
	}
	return check, err
}
func (m sudokuMap) checkCellForZero(cRow, cCol int8) (check bool, err error) {
	outOfBounds := errors.New("out of bounds")
	//Using switch statements instead of if statements because I like switch over if
	switch /*cRow*/ {
	case cRow < 0:
		return true, outOfBounds
	case cRow > 8:
		return true, outOfBounds
	} //Check rows for out of bounds
	switch /*cCol*/ {
	case cCol < 0:
		return true, outOfBounds
	case cCol > 8:
		return true, outOfBounds
	} //Check columns for out of bounds
	switch /*check if this cell is a 0 in mainSudoku*/ {
	case mainSudoku[cRow][cCol] == 0:
		check = false
	default:
		check = true
	}
	return check, err
}
func (m sudokuMap) runAllChecks(cRow, cCol, number int8) (check bool) {
	checkRow := m.checkRows(cRow, number)
	checkColumn := m.checkColumns(cCol, number)
	checkSubRegions, errCheckSubregions := m.checkSubRegions(cRow, cCol, number)
	checkZeroCell, errZeroCell := m.checkCellForZero(cRow, cCol)
	switch {
	case checkRow == true:
		fmt.Println("checkRows returned true")
		check = true
	case checkColumn == true:
		fmt.Println("checkColumns returned true")
		check = true
	case checkSubRegions == true:
		fmt.Println("checkSubRegions returned true")
		fmt.Printf("Error Encountered: %v", errCheckSubregions)
		check = true
	case checkZeroCell == true:
		fmt.Println("checkZeroCell returned true")
		fmt.Printf("Error Encountered: %v", errZeroCell)
		check = true
	default:
		check = false
	}

	return check
}
func (m *sudokuMap) modifyCell(cRow, cCol, number int8) {
	if m.runAllChecks(cRow, cCol, number) == false {
		m[cRow][cCol] = number
		fmt.Println("Valid Move")
	} else {
		fmt.Println("Invalid Move: sudokuMap is not modified")
	}
}

func main() {
	var (
		currentSudoku sudokuMap
	)

	currentSudoku = currentSudoku.newSudoku()

	//Test 1 Test Zero Check
	fmt.Printf("\n(0,0) = %v\n", currentSudoku[0][0])
	fmt.Printf("Starting Test 1\n")
	currentSudoku.modifyCell(0, 0, 1)

	//Test 2 Test Subregion Check
	fmt.Printf("\n\n(0,2) = %v\n", currentSudoku[0][2])
	fmt.Printf("Starting Test 2\n")
	currentSudoku.modifyCell(0, 2, 9)

	//Test 3 Test Column Check
	fmt.Printf("\n\n(0,8) = %v\n", currentSudoku[0][8])
	fmt.Printf("Starting Test 3\n")
	currentSudoku.modifyCell(0, 8, 5)
	/*Cause a subregion not found error. Will need to fix*/

	//Test 4 Row Check
	fmt.Printf("\n\n(4,3) = %v\n", currentSudoku[4][3])
	fmt.Printf("Starting Test 3\n")
	currentSudoku.modifyCell(4, 3, 5)
}
