package main

import (
	"fmt"

	"github.com/Cosiamo/DocuSort/dsx"
)

// The main func for DocuSort serves as testing for dsx,
// as well as, showing examples for GitHub
func main() {
	// opens files
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	// closes files
	defer dsx.CloseXlsx(file)
	// which sheet the method's access
	sheet := "Sheet1"

	var testCol0 []string = []string{"A", "B", "C", "D"}
	var testRow0 []int = []int{2, 3, 4, 5, 6}
	// group data from specific columns
	// dsx.Group(f *excelize.File, sheet string, columns []string, rows []int) chan string
	for test0 := range dsx.Group(file, sheet, testCol0, testRow0) {
		fmt.Println(test0)
	}

	startRow1 := 1
	endRow1 := 30
	// returns int slices within params passed into the function
	// dsx.RangeOfRows(startRow int, endRow int) chan []int
	for test1 := range dsx.RangeOfRows(startRow1, endRow1) {
		fmt.Println(test1)
	}

	startCol2 := "A"
	endCol2 := "BZ"
	// returns string slices within params passed into the function
	// dsx.RangeOfCols(startCol string, endCol string) chan []string
	for test2 := range dsx.RangeOfCols(startCol2, endCol2) {
		fmt.Println(test2)
	}

	// returns all data from the spreadsheet by column
	// dsx.AllDataByCols(f *excelize.File, sheet string) chan string
	for test3 := range dsx.AllDataByCols(file, sheet) {
		fmt.Print(test3)
		fmt.Println()
	}

	// returns all data from the spreadsheet by row
	// dsx.AllDataByRows(f *excelize.File, sheet string) chan string
	for test4 := range dsx.AllDataByRows(file, sheet) {
		fmt.Print(test4)
		fmt.Println()
	}

	var testCol5 []string = []string{"A", "B", "C", "D"}
	var testRow5 []int = []int{2, 3, 4, 5, 6}
	// returns the coordinates of cells from the params passed
	// dsx.MapGroupCoords(columns []string, rows []int) chan map[int]int
	for test5 := range dsx.MapGroupCoords(testCol5, testRow5) {
		fmt.Println(test5)
	}
}