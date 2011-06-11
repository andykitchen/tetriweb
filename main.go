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

func (b *Board) setCell(i, j, k int) {
	b.cells[i*WIDTH + j] = k
}

func (b *Board) rowFull(i int) (r bool) {
	r = true
	
	for j := 0; j < WIDTH; j++ {
		r = r && b.getCell(i, j) > 0
		
		if !r {
			break
		}
	}
	
	return
}

func (b *Board) clearRow(row int) {
	for j := 0; j < WIDTH; j++ {
		b.setCell(row, j, 0)
	}
}

func (b *Board) copyRow(row_from, row_to int) {
	for j := 0; j < WIDTH; j++ {
		cell := b.getCell(row_from, j)
		b.setCell(row_to, j, cell)
	}
}

func (b *Board) removeRow(row int) {
	for i := row; i < HEIGHT-1; i++ {
		b.copyRow(i+1, i)
	}

	b.clearRow(HEIGHT-1)
}


func (b *Board) String() (s string) {
	for i := HEIGHT-1; i >= 0; i-- {
		for j := 0; j < WIDTH; j++ {
			s += strconv.Itoa(b.getCell(i, j))
		}
		
		s += "\n"
	}
	
	return
}

func main() {
	b := new(Board)

	for j := 0; j < WIDTH; j++ {
		b.setCell(2, j, 1)
	}

	b.setCell(10, 3, 1)
	b.setCell(10, 4, 1)
	b.setCell(11, 3, 1)
	b.setCell(11, 5, 1)

	b.setCell(19, 5, 1)
	b.setCell(19, 6, 1)
	b.setCell(19, 7, 1)


	b.copyRow(2, 4)
	b.removeRow(2)
	b.removeRow(2)

	fmt.Printf(b.String())
}
