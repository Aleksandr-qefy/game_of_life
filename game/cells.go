package game

import (
	"bufio"
	"fmt"
)

func InitCells(m, n int) CellsT {
	cells := make([]CellsRowT, m)
	for i := 0; i < m; i++ {
		cells[i] = make([]CellT, n)
	}
	return cells
}

func (c CellsT) NearCount(i, j int) int {
	n := 0
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if (k == 0) && (l == 0) {
				continue
			}
			if c[i+k][j+l] {
				n += 1
			}
		}
	}
	return n
}

func ReadCells(in *bufio.Reader) CellsT {
	var m, n int
	fmt.Fscan(in, &m, &n)

	cells := InitCells(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var char string
			fmt.Fscan(in, &char)
			switch char {
			case "1":
				cells[i][j] = true
			case "0":
				cells[i][j] = false
			}
		}
	}
	return cells
}
