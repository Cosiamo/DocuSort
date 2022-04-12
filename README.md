# DocuSort
DocuSort is a collection of Go packages that are meant to make sorting `.xlsx`, `.csv`, and other spreadsheet files much easier. It currently only supports `.xlsx` files (the dsx package), but I have plans in the future to support various other file extensions.

# Table of Contents
- [dsx](#dsx)
    - [Open and close files](#opening-and-closing-files)
    - [Group columns and rows](#group-columns-and-rows)
		- [Group method](#group-method)
		- [RangeOfRows method](#rangeofrows-method)
		- [RangeOfCols method](#rangeofcols-method)
	- [Map group coordinates](#mapgroupcoords-method)
    - [Get all data by rows](#to-get-all-data-by-row)
    - [Get all data by columns](#to-get-all-data-by-column)

## dsx
Uses the Go package [excelize](https://github.com/qax-os/excelize) to access `.xlsx` files. Dsx has built in methods to make tasks that involve large data sets much simpler.

### Installation
```
go get github.com/Cosiamo/DocuSort/dsx
```
```go
import (
    "github.com/Cosiamo/DocuSort/dsx"
)
```

### Opening and closing files
```go 
func main() {
    // OpenXlsx("[PATH]/<file name>")
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
    // either call CloseXlsx with 'defer'
    // or at the end of the function
	defer dsx.CloseXlsx(file)
}
```

### Group columns and rows
#### Group method
This is for when you need to get specific groups of data. Let's say you're working with a spreadsheet that has list of computers. In column "A" you have the brand names, in column "B" the product number, in column "C" the MAC address, and column "D" if they're active or not. If you want to get all data from a specific brand, the `Group` method will return all the cells with the relevant info.
```go 
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(file)
    sheet := "Sheet1"

    var columns []string = []string{"A", "B", "C", "D"}
    // want to return all info about Apple computers from the spreadsheet
    // all the rows that have the value "Apple" as brand
	var rows []int = []int{2, 3, 4, 5, 6}

	// returns values of cells from an array of strings and ints
	for res := range dsx.Group(f, sheet, columns, rows) {
		fmt.Println(res)
	}
}
```

<img src="/imgs/FuncGroupRes.png">

#### RangeOfRows method
If you're working with dozens, hundreds, or even thousands of rows that you need grouped, the `RangeOfRows` method will let you set a start-row and end-row so you don't have to manually type in all row numbers. (This gives the same result as the previous example)
```go
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(f)
	sheet := "Sheet1"

	var columns []string = []string{"A", "B", "C", "D"}

	startRow := 2
	endRow := 6
	// returns int slices within params passed into the method
	for rows := range dsx.RangeOfRows(startRow, endRow) {
		for res := range dsx.Group(file, sheet, columns, rows) {
			fmt.Println(res)
		}
	}
}
```

#### RangeOfCols method
If you're working a lot of columns, the `RangeOfCols` method will let you set a start-column and end-column. It can take values anywhere from "A" to "ZZ".
```go
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(f)
	sheet := "Sheet1"

	startCol := "A"
	endCol := "D"
	startRow := 2
	endRow := 6
	for rows := range dsx.RangeOfRows(startRow, endRow) {
		// returns string slices within params passed into the method
		for columns := range RangeOfCols(startCol, endCol) {
			for res := range dsx.Group(file, sheet, columns, rows) {
				fmt.Println(res)
			}
		}
	}
}
```

### MapGroupCoords method
Maps the cells coordinates to the parameters that are provided.
```go
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(f)
	sheet := "Sheet1"

	var columns []string = []string{"A", "B", "C", "D"}
	var rows []int = []int{2, 3, 4, 5, 6}
	for res := range dsx.MapGroupCoords(columns, rows) {
		fmt.Println(res)
	}
}
```

<img src="/imgs/MapGroupCoordsRes.png">

### To get all data by row
```go
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(file)
    sheet := "Sheet1"

    // returns all data from the spreadsheet by row
	for data := range dsx.AllDataByRows(file, sheet) {
		fmt.Print(data)
		fmt.Println()
	}
}
```

<img src="/imgs/AllDataByRowsRes.png">

### To get all data by column
```go 
func main() {
	file := dsx.OpenXlsx("spreadsheets/testSheet.xlsx")
	defer dsx.CloseXlsx(file)
    sheet := "Sheet1"

    // returns all data from the spreadsheet by column
    for data := range dsx.AllDataByCols(file, sheet) {
		fmt.Print(data)
		fmt.Println()
	}
}
```

<img src="/imgs/AllDataByColsRes.png">
