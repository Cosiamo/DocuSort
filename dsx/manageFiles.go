package dsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// opens .xlsx files
func OpenXlsx(fileName string) (f *excelize.File) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	return f
}

// closes .xlsx files
func CloseXlsx(f *excelize.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}