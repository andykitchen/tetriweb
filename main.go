package main

import (
	"fmt"
	// "syscall"
)


func NewBoard() (b Board) {
	return
}

func abs(i int) (j int) {
	if i < 0 {
		j = -i
	} else {
		j = i
	}

	return
}

func main() {
	b := NewBoard()

	for j := 3; j < 8; j++ {
		b.setCell(j, 2, 1)
		b.setCell(j, 1, 1)
	}

	b.FillLine(2)
	b.play_shape = NewShape(T_SHAPE)
	b.play_shape.RotateClockwise()
	b.sx = 12
	b.sy = 3

	var s string
	for {
		fmt.Print(b.String())
		fmt.Print(b.sx, "------------------\n")
		// syscall.Sleep(1000000000)
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			fmt.Println(err.String())
			return
		}
		if s == "l" {
			fmt.Println("ROTATE!")
			//b.play_shape.RotateClockwise()
		} else {
			fmt.Println("NOROTATE")
		}

		b.Tick()
	}
}
