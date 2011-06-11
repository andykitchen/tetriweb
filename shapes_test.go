package main

import (
	"testing"
	//"fmt"
	//"strconv"
)

func TestShapeSameAs(t *testing.T) {
	s1 := NewShape(T_SHAPE)
	s2 := NewShape([]int{0, 0, 0, 0})
	result := s1.SameAs(s2)
	if result == true {
		t.Errorf("Invalid Same As length")
	}
}

func TestNewShape(t *testing.T) {
	s := NewShape(T_SHAPE)
	if len(s.CurrentState) != len(T_SHAPE) {
		t.Errorf("Invalid current state")
	}
}
