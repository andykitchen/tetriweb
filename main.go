package main

import (
	"fmt"
	"strconv"
)

const (
	WIDTH  = 10
	HEIGHT = 20
	SHAPE_W = 4
	SHAPE_H = 4
)

type Board struct {
	cells [WIDTH*HEIGHT]int
	sx, sy int
}

func (b *Board) getCell(i, j int) int {
	return b.cells[i*WIDTH + j]
}

// func (b *Board) rowFull(i int) (r bool) {
// 	r = true
// 	
// 	for j := 0; j < WIDTH; j++ {
// 		r &&= b.getCell(i, j) > 0
// 		
// 		if !r {
// 			break
// 		}
// 	}
// 	
// 	return
// }

// func (b *Board) removeRow(i int) (r bool) {
// 	r = true
// 	
// 	for j := 0; j < WIDTH; j++ {
// 		r &&= b.getCell(i, j) > 0
// 		
// 		if !r {
// 			break
// 		}
// 	}
// 	
// 	return
// }


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
