package game

type Playground struct {
	Cells CellsT
	W     int
	H     int
	Shift int
}

type CellsT []CellsRowT
type CellsRowT []CellT
type CellT bool
