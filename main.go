package main

import (
	"os"
	"fmt"
	"rand"
	"exec"
	// "syscall"
)

func abs(i int) (j int) {
	if i < 0 {
		j = -i
	} else {
		j = i
	}

	return
}

func PlayTerminal() {
	b := NewBoard()

	// for j := 3; j < 8; j++ {
	// 	b.setCell(j, 2, 1)
	// 	b.setCell(j, 1, 1)
	// }
	// 
	// b.FillLine(2)
	b.play_shape = NewShape(T_SHAPE)
	b.play_shape.RotateClockwise()
	b.sx = 15
	b.sy = 6

	ticks := 0
	rand.Seed(1)
	exec.Run("/bin/stty", []string{"stty", "-icanon", "min", "1", "-echo"},
		nil, "", exec.PassThrough, exec.PassThrough, exec.PassThrough)
	for {

		fmt.Print(b.String())
		fmt.Print(ticks, "------------------\n")
		// syscall.Sleep(1000000000)
		buf := make([]byte, 1)
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println(err.String())
			return
		}

		s := string(buf)

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
