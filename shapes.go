package main

import (
	//"fmt"
	"strconv"
)

var T_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 1, 1, 1, 0,
	0, 0, 0, 0, 0,
	0, 0, 0, 0, 0}

var BLOCK_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 1, 1, 0, 0,
	0, 1, 1, 0, 0,
	0, 0, 0, 0, 0,
	0, 0, 0, 0, 0}

var I_SHAPE []int = []int{
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 0, 0, 0}

var RL_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 1, 0,
	0, 0, 0, 0, 0}

var LL_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 1, 1, 0, 0,
	0, 0, 0, 0, 0}

var S_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 0, 1, 1, 0,
	0, 1, 1, 0, 0,
	0, 0, 0, 0, 0,
	0, 0, 0, 0, 0}

var Z_SHAPE []int = []int{
	0, 0, 0, 0, 0,
	0, 1, 1, 0, 0,
	0, 0, 1, 1, 0,
	0, 0, 0, 0, 0,
	0, 0, 0, 0, 0}

var SHAPES [][]int = [][]int{
	T_SHAPE, BLOCK_SHAPE, I_SHAPE,
	RL_SHAPE, LL_SHAPE, S_SHAPE, Z_SHAPE}

const SHAPE_WIDTH = 5
const SHAPE_HEIGHT = 5
const SHAPE_X_OFFSET = -2
const SHAPE_Y_OFFSET = -2

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Shape struct {
	baseState   []int
	Orientation int
  Colour      int
}

func (s *Shape) CurrentState() (result []int) {
	result = make([]int, 25)

	for i := 0; i < len(result); i++ {
		x := i % SHAPE_WIDTH
		y := (i / SHAPE_WIDTH)

		//fmt.Printf(strconv.Itoa(x) + ", " + strconv.Itoa(y) + "\n")
		result[i] = s.GetCell(SHAPE_X_OFFSET+x,
			SHAPE_Y_OFFSET+y)
	}

	return
}

func (s *Shape) String() (res string) {
	currentState := s.CurrentState()

	for i := 0; i < SHAPE_HEIGHT; i++ {
		for j := 0; j < SHAPE_WIDTH; j++ {
			res += strconv.Itoa(currentState[i*SHAPE_WIDTH+j])
		}

		res += "\n"
	}

	return

}

func (s *Shape) GetCell(x int, y int) (k int) {
	switch s.Orientation {
	case UP:
		k = s.getShapeCell(x, y)
	case RIGHT:
		k = s.getShapeCell(y, -x)
	case DOWN:
		k = s.getShapeCell(x, -y)
	case LEFT:
		k = s.getShapeCell(-y, x)
	}

	return
}

func (s *Shape) getShapeCell(x, y int) int {
	//TODO: add bounds checking
	if x < -2 || x > 2 {
		return 0
	}
	if y < -2 || y > 2 {
		return 0
	}
	index := (x + 2) + (y+2)*SHAPE_WIDTH
	//fmt.Printf(strconv.Itoa(index) + "\n")
	return s.baseState[index]
}

func (s *Shape) RotateClockwise() {
	if s.Orientation == LEFT {
		s.Orientation = UP
	} else {
		s.Orientation += 1
	}
}

func (s *Shape) RotateCounterClockwise() {
	if s.Orientation == UP {
		s.Orientation = LEFT
	} else {
		s.Orientation -= 1
	}
}

func (s1 *Shape) SameAs(s2 *Shape) (result bool) {
	if len(s1.baseState) != len(s2.baseState) {
		return false
	}

	if len(s1.CurrentState()) != len(s2.CurrentState()) {
		return false
	}

	result = true
	for i := 0; i < len(s1.CurrentState()); i++ {
		if s1.CurrentState()[i] != s2.CurrentState()[i] {
			result = false
		}
	}
	return
}

func NewShape(shape []int) (s *Shape) {
	s = new(Shape)
	s.baseState = shape
	return
}
