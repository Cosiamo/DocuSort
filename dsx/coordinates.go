package dsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// returns a map of coordinates of cells from an array of strings and ints
func MapGroupCoords(columns []string, rows []int) (chan map[int]int) {
	ch := make(chan map[int]int)
	go func() {
		// ranges over rows
		for _, r := range rows {
			// ranges over columns
			for _, c := range columns {
				format := fmt.Sprintf("%s%d", c, r)
				i1, i2, err := excelize.CellNameToCoordinates(format)
				if err != nil {
					fmt.Println(err)
					return
				}
				var coord = map[int]int {
					i1: i2,
				}
				ch <- coord
			}
		}
		close(ch)
	}()
	return ch
}