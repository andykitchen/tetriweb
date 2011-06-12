package main

import (
	"strconv"
	"rand"
)

const (
	WIDTH         = 10
	HEIGHT        = 20
	SHAPE_W       = 4
	SHAPE_H       = 4
	TICKS_TO_GRAV = 3
)

type Board struct {
	cells      [WIDTH * HEIGHT]int
	sx, sy     int
	play_shape *Shape
	next_shape *Shape
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

func (b *Board) FillLine(y int) {
	for j := 0; j < WIDTH; j++ {
		b.setCell(y, j, 1)
	}
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

func CheckShapeOverlap(b *Board, newX int, newY int) bool {
	for i := -1; i < HEIGHT+1; i++ {
		for j := -1; j < WIDTH+1; j++ {
			ox, oy := newX-i, newY-j

			if abs(ox) <= 2 && abs(oy) <= 2 {
				outside_board := i < 0 || i >= HEIGHT || j < 0 || j >= WIDTH
				shape_filled  := b.play_shape.GetCell(ox, oy) > 0
				
				if outside_board && shape_filled {
					return true
				} else {
					cell_filled  := b.getCellRaw(i, j) > 0
					if cell_filled && shape_filled {
						return true
					}
				}
			}

		}
	}

	return false
}

func (b *Board) String() (s string) {
	for i := HEIGHT - 1; i >= 0; i-- {
		for j := 0; j < WIDTH; j++ {
			s += strconv.Itoa(b.getCell(i, j))
		}

		s += "\n"
	}

	return
}

var _ticks_betweent_grav int = 0

func (b *Board) Tick() {

	if _ticks_betweent_grav == 3 {
		b.MoveDown()
		_ticks_betweent_grav = 0
	} else {
		_ticks_betweent_grav++
	}

	nextX := b.sx - 1
	if CheckShapeOverlap(b, nextX, b.sy) {
		b.updateShapeToBoard()

		b.play_shape = getRandomShape()
		b.sx = HEIGHT
		b.sy = WIDTH / 2
	}
}

func (b *Board) MoveLeft() {
	nextY := b.sy - 1
	if !CheckShapeOverlap(b, b.sx, nextY) {
		b.sy = nextY
	}
}

func (b *Board) MoveRight() {
	nextY := b.sy + 1
	if !CheckShapeOverlap(b, b.sx, nextY) {
		b.sy = nextY
	}
}

func (b *Board) MoveDown() {
	nextX := b.sx - 1
	if !CheckShapeOverlap(b, nextX, b.sy) {
		b.sx = nextX
	}
}

func (b *Board) updateShapeToBoard() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			ox, oy := b.sx-i, b.sy-j

			if abs(ox) <= 2 && abs(oy) <= 2 {
				k := b.play_shape.GetCell(ox, oy)

				if k > 0 {
					b.setCell(i, j, 1)
				}
			}
		}
	}
}

func (b *Board) RotateClockwise() {
	b.play_shape.RotateClockwise()

	if CheckShapeOverlap(b, b.sx, b.sy) {
		//undo
		b.play_shape.RotateCounterClockwise()
	}
}

func (b *Board) RotateCounterClockwise() {
	b.play_shape.RotateCounterClockwise()

	if CheckShapeOverlap(b, b.sx, b.sy) {
		//Undo
		b.play_shape.RotateClockwise()
	}
}

func getRandomShape() (shape *Shape) {
	num_shapes := len(SHAPES)
	index := rand.Intn(num_shapes)
	shape = NewShape(SHAPES[index])

	new_orient := rand.Intn(LEFT)
	shape.Orientation = new_orient
	return
}
