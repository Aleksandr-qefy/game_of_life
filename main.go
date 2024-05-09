package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// rowsn, colsn - sizes of grid
const (
	rowsn = 18
	colsn = 18
)

type Playground struct {
	cells Cells
	w     int
	h     int
	shift int
}

type Cells []CellsRow
type CellsRow []Cell
type Cell bool

func initCells(m, n int) Cells {
	cells := make([]CellsRow, m)
	for i := 0; i < m; i++ {
		cells[i] = make([]Cell, n)
	}
	return cells
}

func initPlayground(m, n, shift int) Playground {
	playground := Playground{}
	playground.cells = initCells(m, n)
	playground.w = n
	playground.h = m
	playground.shift = shift
	return playground
}

func (pg Playground) show() {
	fmt.Print("\033[2J\033[H")
	for i := pg.shift; i+pg.shift < pg.h; i++ {
		for j := pg.shift; j+pg.shift < pg.w; j++ {
			switch pg.cells[i][j] {
			case true:
				fmt.Print("X")
			case false:
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func (pg Playground) addFigure(i, j, h, w int, cells Cells) {
	for k := 0; k < h; k++ {
		for l := 0; l < w; l++ {
			pg.cells[i+k][j+l] = cells[k][l]
		}
	}
}

func (pg Playground) copy() Playground {
	pg2 := initPlayground(pg.w, pg.h, pg.shift)
	for i := 0; i < pg.h; i++ {
		for j := 0; j < pg.w; j++ {
			pg2.cells[i][j] = pg.cells[i][j]
		}
	}
	return pg2
}

func (c Cells) nearCount(i, j int) int {
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

func (pg Playground) tick() {
	pg2 := pg.copy()
	for i := 1; i+1 < pg2.h; i++ {
		for j := 1; j+1 < pg2.w; j++ {
			near := pg2.cells.nearCount(i, j)
			if pg2.cells[i][j] {
				pg.cells[i][j] = (near == 2) || (near == 3)
			} else {
				pg.cells[i][j] = near == 3
			}
		}
	}
}

func readCells(in *bufio.Reader) Cells {
	var m, n int
	fmt.Fscan(in, &m, &n)

	cells := initCells(m, n)
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

func (pg Playground) loadFigure(in *bufio.Reader) {
	var i, j int
	fmt.Fscan(in, &i, &j)

	figure := readCells(in)
	pg.addFigure(i, j, len(figure), len(figure[0]), figure)
}

func main() {
	if len(os.Args) != 2 {
		panic("correct way to run program:\n    go run main.go figure.txt")
	}

	figureFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer figureFile.Close()

	in := bufio.NewReader(figureFile)

	playground := initPlayground(rowsn+2, colsn+2, 1)
	playground.loadFigure(in)
	for {
		playground.show()
		time.Sleep(1 * time.Second)
		playground.tick()
	}
}
