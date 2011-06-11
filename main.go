package main

import (
	"fmt"
	"strconv"
)

const (
	WIDTH  = 10
	HEIGHT = 20
)

type Board struct {
	cells [WIDTH*HEIGHT]int
}

func (b *Board) getCell(i, j int) int {
	return b.cells[i*WIDTH + j]
}

func (b *Board) String() (s string) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			s += strconv.Itoa(b.getCell(i, j))
		}
		
		s += "\n"
	}
	
	return
}

func main() {
	b := new(Board)
	fmt.Printf(b.String())
}
