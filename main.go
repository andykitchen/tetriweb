package main

import (
	"fmt"
	"strconv"
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

	for j := 0; j < WIDTH; j++ {
		b.setCell(2, j, 1)
	}

	b.play_shape = NewShape(T_SHAPE)

	b.sx = 10
	b.sy = 5

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
			b.play_shape.RotateClockwise()
		} else {
			fmt.Println("NOROTATE")
		}

		b.Tick()
	}
}
