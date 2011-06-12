package main

import (
	"testing"
	//"fmt"
	//"strconv"
)

func TestBoard(t *testing.T) {
	b := new(Board)
	s := b.String()

	if s != "0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n" {
		t.Errorf("invalid base state")
	}
}

func TestRowFull(t *testing.T) {
	b := new(Board)

	for j := 0; j < WIDTH; j++ {
		b.setCell(2, j, 1)
	}

	if !b.checkRowFull(2) {
		t.Errorf("row was full")
	}

	if b.checkRowFull(3) {
		t.Errorf("row wasn't full")
	}
}

func TestPlacingShape(t *testing.T) {
	b := new(Board)

	b.play_shape = NewShape(T_SHAPE)
	b.sx = 10
	b.sy = 5

	s := b.String()

	if s != "0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000010000\n0000011000\n0000010000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n" {
		t.Errorf("invalid new play state")
	}
}

func TestMoveShapeLeft(t *testing.T) {
	b := new(Board)

	b.play_shape = NewShape(T_SHAPE)
	b.sx = 10
	b.sy = 5

	b.MoveLeft()
	s := b.String()
	if s != "0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000100000\n0000110000\n0000100000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n" {
		t.Errorf("invalid move left")
	}
}

func TestMoveShapeRight(t *testing.T) {
	b := new(Board)

	b.play_shape = NewShape(T_SHAPE)
	b.sx = 10
	b.sy = 5
	b.MoveRight()
	s := b.String()

	if s != "0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000001000\n0000001100\n0000001000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n" {
		t.Errorf("invalid move right")
	}
}

func TestMoveShapeDown(t *testing.T) {
	b := new(Board)

	b.play_shape = NewShape(T_SHAPE)
	b.sx = 10
	b.sy = 5
	b.MoveDown()
	s := b.String()

	if s != "0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000010000\n0000011000\n0000010000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n0000000000\n" {
		t.Errorf("invalid move down")
	}
}
