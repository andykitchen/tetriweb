package main

import (
	"testing"
	"fmt"
	//"strconv"
)

func TestShapeSameAs(t *testing.T) {
	s1 := NewShape(T_SHAPE)
	s2 := NewShape([]int{0, 0, 0, 0})

	result := s1.SameAs(s2)
	if result == true {
		t.Errorf("Invalid Same As length")
	}

  s3 := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 1, 1, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 0, 0, 0})

	result = s1.SameAs(s3)
	if result == true {
		t.Errorf("Invalid Same As current state")
	}

  s4 := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 1, 1, 1, 0,
	  0, 0, 0, 0, 0,
	  0, 0, 0, 0, 0})

	result = s1.SameAs(s4)
	if result != true {
		t.Errorf("Invalid Same As current state")
	}
}

func TestNewShape(t *testing.T) {
	s := NewShape(T_SHAPE)
	if len(s.CurrentState()) != len(T_SHAPE) {
		t.Errorf("Invalid current state")
	}
}

func TestRotateClockwiseRIGHT(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 1, 1, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateClockwise()
  
  if s.Orientation != RIGHT {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

func TestRotateClockwiseDOWN(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 0, 0, 0,
	  0, 1, 1, 1, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateClockwise()
  s.RotateClockwise()
  
  if s.Orientation != DOWN {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

func TestRotateClockwiseLEFT(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 1, 1, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateClockwise()
  s.RotateClockwise()
  s.RotateClockwise()
  
  if s.Orientation != LEFT {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

func TestRotateClockwiseFULLCIRCLE(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 1, 1, 1, 0,
	  0, 0, 0, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateClockwise()
  s.RotateClockwise()
  s.RotateClockwise()
  s.RotateClockwise()
  
  if s.Orientation != UP {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

func TestRotateCounterClockwiseLEFT(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 1, 1, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateCounterClockwise()
  
  if s.Orientation != LEFT {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

func TestRotateCounterClockwiseFULLCIRCLE(t *testing.T) {
  expected := NewShape([]int{
    0, 0, 0, 0, 0,
	  0, 0, 1, 0, 0,
	  0, 1, 1, 1, 0,
	  0, 0, 0, 0, 0,
	  0, 0, 0, 0, 0})

  s := NewShape(T_SHAPE)

  s.RotateCounterClockwise()
  s.RotateCounterClockwise()
  s.RotateCounterClockwise()
  s.RotateCounterClockwise()
  
  if s.Orientation != UP {
    t.Errorf("Invalid orientation")
  }
  
  fmt.Printf(s.String() + "\n")

  result := s.SameAs(expected)
  if result != true {
    fmt.Printf(s.String())
    t.Errorf("Invalid rotate clockwise")
  }
}

