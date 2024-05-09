package main

import (
	"bufio"
	"fmt"
	gm "game/game"
	"os"
	"time"
)

// rowsn, colsn - sizes of grid
const (
	rowsn = 18
	colsn = 18
)

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

	playground := gm.InitPlayground(rowsn+2, colsn+2, 1)
	playground.LoadFigure(in)
	for {
		playground.Show()
		time.Sleep(1 * time.Second)
		playground.Tick()
	}
}
