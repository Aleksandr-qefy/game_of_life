package game

import (
	"bufio"
	"fmt"
)

func InitPlayground(m, n, shift int) Playground {
	playground := Playground{}
	playground.Cells = InitCells(m, n)
	playground.W = n
	playground.H = m
	playground.Shift = shift
	return playground
}

func (pg Playground) Show() {
	fmt.Print("\033[2J\033[H")
	for i := pg.Shift; i+pg.Shift < pg.H; i++ {
		for j := pg.Shift; j+pg.Shift < pg.W; j++ {
			switch pg.Cells[i][j] {
			case true:
				fmt.Print("X")
			case false:
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func (pg Playground) AddFigure(i, j, h, w int, cells CellsT) {
	for k := 0; k < h; k++ {
		for l := 0; l < w; l++ {
			pg.Cells[i+k][j+l] = cells[k][l]
		}
	}
}

func (pg Playground) Copy() Playground {
	pg2 := InitPlayground(pg.W, pg.H, pg.Shift)
	for i := 0; i < pg.H; i++ {
		for j := 0; j < pg.W; j++ {
			pg2.Cells[i][j] = pg.Cells[i][j]
		}
	}
	return pg2
}

func (pg Playground) Tick() {
	pg2 := pg.Copy()
	for i := 1; i+1 < pg2.H; i++ {
		for j := 1; j+1 < pg2.W; j++ {
			near := pg2.Cells.NearCount(i, j)
			if pg2.Cells[i][j] {
				pg.Cells[i][j] = (near == 2) || (near == 3)
			} else {
				pg.Cells[i][j] = near == 3
			}
		}
	}
}

func (pg Playground) LoadFigure(in *bufio.Reader) {
	var i, j int
	fmt.Fscan(in, &i, &j)

	figure := ReadCells(in)
	pg.AddFigure(i, j, len(figure), len(figure[0]), figure)
}
