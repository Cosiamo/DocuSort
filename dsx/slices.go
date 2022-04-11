package dsx

func RangeOfRows(startRow int, endRow int) chan []int {
	ch := make(chan []int)
	i := startRow
	go func() {
		for i <= endRow {
			ch <- []int{i}
			i++
		}
		close(ch)
	}()
	return ch
}