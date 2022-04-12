package dsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// opens files in .xlsx format
func OpenXlsx(fileName string) (f *excelize.File) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	return f
}

// closes files in .xlsx format
func CloseXlsx(f *excelize.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}