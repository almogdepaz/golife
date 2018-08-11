package golife

import (
	"testing"
)

func TestSet(t *testing.T) {
	f := NewField(3, 3)
	f.Set(0, 0, true)
	if !f.s[0][0] {
		t.Error("Set failed to change value")
	}
}

func TestAlive(t *testing.T) {
	f := NewField(3, 3)
	f.Set(0, 0, true)
	if !f.Alive(0, 0) {
		t.Error("Alive failed to retrive correct value")
	}
}

func TestACountLiveNeighbors(t *testing.T) {
	f := NewField(3, 3)
	f.Set(0, 0, true)
	f.Set(0, 1, true)
	if f.CountLiveNeighbors(1, 0) != 2 {
		t.Error("CountLiveNeighbors failed to calculate correct value")
	}
}
