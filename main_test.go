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

  if !b.rowFull(2) {
    t.Errorf("row was full")
  }

  if b.rowFull(3) {
    t.Errorf("row wasn't full")
  }
}

