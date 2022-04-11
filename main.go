package main

import (
	"fmt"

	"github.com/Cosiamo/DocuSort/dsx"
)

func main() {
	f := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(f)
	sheet := "Sheet1"

	// var testCol []string = []string{"A", "B", "C", "D"}
	// // var testRow []int = []int{2, 3, 4, 5, 6}

	// startRow := 2
	// endRow := 6
	// // returns int slices within params passed into the function
	// for testRow := range dsx.RangeOfRows(startRow, endRow) {
	// 	// returns values of cells from an array of strings and ints
	// 	for test0 := range dsx.Group(f, sheet, testCol, testRow) {
	// 		fmt.Println(test0)
	// 	}
	// }

	

	// // returns all data from the spreadsheet by column
	// for test2 := range dsx.AllDataByCols(f, sheet) {
	// 	fmt.Print(test2)
	// 	fmt.Println()
	// }

	// returns all data from the spreadsheet by row
	for test3 := range dsx.AllDataByRows(f, sheet) {
		fmt.Print(test3)
		fmt.Println()
	}

}