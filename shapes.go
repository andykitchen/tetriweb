package main

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

type Shape struct {
	CurrentState []int
}

func (s1 *Shape) SameAs(s2 *Shape) (result bool) {
	if len(s1.CurrentState) != len(s2.CurrentState) {
		return false
	}

	result = true
	for i := 0; i < len(s1.CurrentState); i++ {
		if s1.CurrentState[i] != s2.CurrentState[i] {
			result = false
		}
	}
	return
}

func NewShape(shape []int) (s *Shape) {
	s = new(Shape)
	s.CurrentState = shape
	return
}
