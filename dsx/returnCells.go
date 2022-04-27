package dsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// returns values of cells from an array of strings and ints
//
// for test0 := range dsx.Group(file, sheet, columns, rows) {
// 	fmt.Println(test0)
// }
func Group(f *excelize.File, sheet string, columns []string, rows []int) (chan string) {
	cellVals := make(chan string)
	go func() {
		// ranges over rows
		for _, r := range rows {
			// ranges over columns
			for _, c := range columns {
				// formats values from columns and rows into string
				format := fmt.Sprintf("%s%d", c, r)
				// retrieves info from cells with corresponding column and row
				cellVal, err := f.GetCellValue(sheet, format)
				if err != nil {
					fmt.Println(err)
					return
				}
				cellVal = append(cellVal, " ")
				cellVals <- cellVal
			}
		}
		close(cellVals)
	}()
	return cellVals
}

// returns all data from the spreadsheet by column
func AllDataByCols(f *excelize.File, sheet string) (chan string) {
	ch := make(chan string)
	go func() {
		cols, err := f.Cols(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		for cols.Next() {
			col, err := cols.Rows()
			if err != nil {
				fmt.Println(err)
			}
			for _, rowCell := range col {
				s := fmt.Sprint(rowCell)
				ch <- s
			}
		}
		close(ch)
	}()
	return ch
}

// returns all data from the spreadsheet by row
func AllDataByRows(f *excelize.File, sheet string) (chan string) {
	ch := make (chan string)
	go func() {
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				fmt.Println(err)
			}
			for _, colCell := range row {
				s := fmt.Sprint(colCell)
				ch <- s
			}
		}
		close(ch)
	}()
	return ch
}