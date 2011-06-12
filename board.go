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

const (
  BLANK = iota
  PLAYING
  FINISHED
)

type Board struct {
	cells      [WIDTH * HEIGHT]int
	sx, sy     int
	play_shape *Shape
	next_shape *Shape
  state      int
  _ticks_betweent_grav int
}

func NewBoard() (b Board) {
	return
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
				shape_filled := b.play_shape.GetCell(ox, oy) > 0

				if outside_board && shape_filled {
					return true
				}

				if !outside_board {
					cell_filled := b.getCellRaw(i, j) > 0
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

func (b *Board) Encode(player_id int) (data []uint8) {
	data = make([]uint8, WIDTH*HEIGHT+1)
	var n int
  data[0] = uint8(player_id)
	n = 1
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			k := b.getCell(i, j)
			data[n] = strconv.Itoa(k)[0]
			n++
		}
	}

	return
}



func (b *Board) Tick() {
  if b.state == FINISHED {return}
  b.state = PLAYING

	if b._ticks_betweent_grav == 5 {
		b.MoveDown()
		b._ticks_betweent_grav = 0
	} else {
		b._ticks_betweent_grav++
	}

	nextX := b.sx - 1
	if CheckShapeOverlap(b, nextX, b.sy) {
		b.updateShapeToBoard()

    n := b.AddShape()
    if n > 0 {b.state = FINISHED}
	}
	
	for row := HEIGHT-1; row >= 0; row-- {
		if b.checkRowFull(row) {
			b.removeRow(row)
		}
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

func (b *Board) GetBoardState() []int {
	result := make([]int, WIDTH*HEIGHT)
	for i := HEIGHT - 1; i >= 0; i-- {
		for j := 0; j < WIDTH; j++ {
			result[i*WIDTH+j] = b.getCell(i, j)
		}
	}
	return result
}

func (b *Board) AddShape() (n int) {
  b.play_shape = getRandomShape()
  startX := HEIGHT - 2
  b.sy = WIDTH / 2
  n = 0
  for {
    if !CheckShapeOverlap(b, startX, b.sy) {
      break
    }
    startX--
    n ++ 
  }
  b.sx = startX
  return n
}

func (b *Board) DropShape() {
  i := 0
  nextX := 0
  for {
    nextX = b.sx - i
	  if CheckShapeOverlap(b, nextX, b.sy) {
		  i -= 1
      break;
	  }
    i++
  }
  b.sx -= i
}
