package main

import (
	"strconv"
)

const (
	WIDTH   = 10
	HEIGHT  = 20
	SHAPE_W = 4
	SHAPE_H = 4
)

type Board struct {
	cells      [WIDTH * HEIGHT]int
	sx, sy     int
	play_shape *Shape
}

func (b *Board) getCell(i, j int) (k int) {
	if b.play_shape != nil {
		ox, oy := b.sx-i, b.sy-j
		if abs(ox) <= 2 && abs(oy) <= 2 {
			k = b.play_shape.GetCell(ox, oy)
			if k > 0 {
				return
			}
		}
	}

	k = b.getCellRaw(i, j)
	return
}

func (b *Board) getCellRaw(i, j int) int {
	return b.cells[i*WIDTH+j]
}

func (b *Board) setCell(i, j, k int) {
	b.cells[i*WIDTH+j] = k
}

func (b *Board) checkRowFull(i int) (r bool) {
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
		cell := b.getCellRaw(row_from, j)
		b.setCell(row_to, j, cell)
	}
}

func (b *Board) removeRow(row int) {
	for i := row; i < HEIGHT-1; i++ {
		b.copyRow(i+1, i)
	}

	b.clearRow(HEIGHT - 1)
}

// func (b *Board) checkShapeCollision() bool {
// 	for i = 0; i < HEIGHT; i++ {
// 		for j = 0; j < WIDTH; j++ {
// 		ox, oy := b.sx - i, b.sy - j;
// 		if abs(ox) <= 2 && abs(oy) <= 2 {
// 			k = b.play_shape.GetCell(ox, oy)
// 			if k > 0 && b.getCellRaw(i, j) > 0 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func (b *Board) String() (s string) {
	for i := HEIGHT - 1; i >= 0; i-- {
		for j := 0; j < WIDTH; j++ {
			s += strconv.Itoa(b.getCell(i, j))
		}

		s += "\n"
	}

	return
}

func (b *Board) Tick() {
	b.sx -= 1
}

func (b *Board) MoveLeft() {
	b.sy -= 1
}

func (b *Board) MoveRight() {
	b.sy += 1
}

func (b *Board) MoveDown() {
	b.sx -= 1
}
