package main

import (
	"fmt"
	"rand"
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
	b.sx = 5
	b.sy = 6

	ticks := 0
	rand.Seed(1)
	var s string
	for {

		fmt.Print(b.String())
		fmt.Print(ticks, "------------------\n")
		// syscall.Sleep(1000000000)
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			fmt.Println(err.String())
			return
		}

		switch s {
		case "h":
			fmt.Println("LEFT!")
			b.MoveLeft()
		case "l":
			fmt.Println("RIGHT!")
			b.MoveRight()
		case "j":
			fmt.Println("DOWN!")
			b.MoveDown()
		case "u":
			fmt.Println("ROTATE!")
			b.RotateCounterClockwise()
		case "o":
			fmt.Println("ROTATE!")
			b.RotateClockwise()
		}

		b.Tick()
		ticks++
	}
}
