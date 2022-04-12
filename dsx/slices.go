package dsx

import (
	"fmt"
	"os"
	"sync"
)

var colMapStr = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

var colMapInt = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

var (
	z = 26
	aa = 27
	az = 52
	ba = 53
	bz = 78
	ca = 79
	cz = 104
	da = 105
	dz = 130
	ea = 131
	ez = 156
	fa = 157
	fz = 182
	ga = 183
	gz = 208
	ha = 209
	hz = 234
	ia = 235
	iz = 260
	ja = 261
	jz = 286
	ka = 287
	kz = 312
	la = 313
	lz = 338
	ma = 339
	mz = 364
	na = 365
	nz = 390
	oa = 391
	oz = 416
	pa = 417
	pz = 442
	qa = 443
	qz = 468
	ra = 469
	rz = 494
	sa = 495
	sz = 520
	ta = 521
	tz = 546
	ua = 547
	uz = 572
	va = 573
	vz = 598
	wa = 599
	wz = 624
	xa = 625
	xz = 650
	ya = 651
	yz = 676
	za = 677
	zz = 702
)

// returns a range of int slices
//
// for res := range dsx.RangeOfRows(startRow, endRow) {
// 	fmt.Println(res)
// }
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

// returns a range of string slices
// 
// for res := range dsx.RangeOfCols(startCol, endCol) {
// 	fmt.Println(res)
// }
func RangeOfCols(startCol string, endCol string) (chan []string) {
	ch := make(chan []string)
	var wg sync.WaitGroup
	go func() {
		// if range is anywhere from "A" to "Z"
		if len(startCol) == 1 && len(endCol) == 1 {
			// finds the int associated with the first column
			start, foundStart := colMapInt[startCol]
			if !foundStart {
				fmt.Println("Sorry, err with first entry")
			}
	
			// finds the int associated with the second column
			end, foundEnd := colMapInt[endCol]
			if !foundEnd {
				fmt.Println("Sorry, err with second entry")
			}
	
			if end < start {
				fmt.Println("Variables were passed in the wrong order in RangeOfCols")
				os.Exit(1)
			}

			// loop everytime the first int is smaller than the second
			for start <= end {
				s := fmt.Sprint(colMapStr[start])
				ch <- []string{s}
				start++
			}
		// if range is anywhere from "A" to "ZZ"
		} else if len(startCol) == 1 && len(endCol) == 2 {
			// finds the int associated with the first column
			start, foundStart := colMapInt[startCol]
			if !foundStart {
				fmt.Println("Sorry, err with first entry")
			}
	
			// index 0 of the second column passed to method
			index0 := string(endCol[0])
			// finds the int associated with index 0
			in0, foundEnd0 := colMapInt[index0]
			if !foundEnd0 {
				fmt.Println("Sorry, err with second entry")
			}
	
			// index 1 of the second column passed to method
			index1 := string(endCol[1])
			// finds the int associated with index 1
			in1, foundEnd1 := colMapInt[index1]
			if !foundEnd1 {
				fmt.Println("Sorry, err with second entry")
			}
	
			colBreak := (in0 * 26)
			end := colBreak + in1
	
			go MulStringValLoop(&wg, ch, start, end)
			wg.Add(1)
			wg.Wait()
		// if range is anywhere from "AA" to "ZZ"
		} else if len(startCol) == 2 && len(endCol) == 2 {
			strIndex0 := string(startCol[0])
			startIndex0, foundStart0 := colMapInt[strIndex0]
			if !foundStart0 {
				fmt.Println("Sorry, err with first entry")
			}

			strIndex1 := string(startCol[1])
			startIndex1, foundStart1 := colMapInt[strIndex1]
			if !foundStart1 {
				fmt.Println("Sorry, err with first entry")
			}

			// index 0 of the second column passed to method
			index0 := string(endCol[0])
			// finds the int associated with index 0
			in0, foundEnd0 := colMapInt[index0]
			if !foundEnd0 {
				fmt.Println("Sorry, err with second entry")
			}
	
			// index 1 of the second column passed to method
			index1 := string(endCol[1])
			// finds the int associated with index 1
			in1, foundEnd1 := colMapInt[index1]
			if !foundEnd1 {
				fmt.Println("Sorry, err with second entry")
			}

			start := (startIndex0 * 26) + startIndex1

			colBreak := (in0 * 26)
			end := colBreak + in1
			
			if end < start {
				fmt.Println("Variables were passed in the wrong order in RangeOfCols")
				os.Exit(1)
			}
			
			go MulStringValLoop(&wg, ch, start, end)
			wg.Add(1)
			wg.Wait()
		} else {
			fmt.Println("RangeOfCols only accepts values anywhere from 'A' to 'ZZ'")
			os.Exit(1)
		}
		close(ch)
	}()
	return ch
}

// loops string slices within paramters of the global
// variables that were mapped in the slices.go file
func MulStringValLoop(wg *sync.WaitGroup, ch chan []string, start int, end int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i <= z {
			s := fmt.Sprint(colMapStr[i])
			ch <- []string{s}
		}
		if i >= aa && i <= az {
			newInd1 := i - z
			newInd0 := 1
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ba && i <= bz {
			newInd1 := i - az
			newInd0 := 2
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ca && i <= cz {
			newInd1 := i - bz
			newInd0 := 3
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= da && i <= dz {
			newInd1 := i - cz
			newInd0 := 4
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ea && i <= ez {
			newInd1 := i - dz
			newInd0 := 5
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= fa && i <= fz {
			newInd1 := i - ez
			newInd0 := 6
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ga && i <= gz {
			newInd1 := i - fz
			newInd0 := 7
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ha && i <= hz {
			newInd1 := i - gz
			newInd0 := 8
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ia && i <= iz {
			newInd1 := i - hz
			newInd0 := 9
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ja && i <= jz {
			newInd1 := i - iz
			newInd0 := 10
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ka && i <= kz {
			newInd1 := i - jz
			newInd0 := 11
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= la && i <= lz {
			newInd1 := i - kz
			newInd0 := 12
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ma && i <= mz {
			newInd1 := i - lz
			newInd0 := 13
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= na && i <= nz {
			newInd1 := i - mz
			newInd0 := 14
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= oa && i <= oz {
			newInd1 := i - nz
			newInd0 := 15
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= pa && i <= pz {
			newInd1 := i - oz
			newInd0 := 16
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= qa && i <= qz {
			newInd1 := i - pz
			newInd0 := 17
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ra && i <= rz {
			newInd1 := i - qz
			newInd0 := 18
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= sa && i <= sz {
			newInd1 := i - rz
			newInd0 := 19
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ta && i <= tz {
			newInd1 := i - sz
			newInd0 := 20
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ua && i <= uz {
			newInd1 := i - tz
			newInd0 := 21
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= va && i <= vz {
			newInd1 := i - uz
			newInd0 := 22
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= wa && i <= wz {
			newInd1 := i - vz
			newInd0 := 23
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= xa && i <= xz {
			newInd1 := i - wz
			newInd0 := 24
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= ya && i <= yz {
			newInd1 := i - xz
			newInd0 := 25
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
		}
		if i >= za && i <= zz {
			newInd1 := i - yz
			newInd0 := 26
			t := fmt.Sprint(colMapStr[newInd0], colMapStr[newInd1])
			ch <- []string{t}
			
		}
	}
}