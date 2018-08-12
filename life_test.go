package golife

import (
	"testing"
	"strings"
)

var (
	input           = []bool{false, false, false, true, true, true, false, true, false}
	expectedStr     = "0 0 0 1 1 1 0 1 0"
	stepInput       = []bool{false, false, false, true, false, false, true, false, true,}
	expectedStepStr = "0 0 0 0 1 0 0 1 0"
)

func TestString(t *testing.T) {
	l := NewLife(3, 3, input, Next)
	str := l.String()
	if strings.Compare(str, expectedStr) != 0 {
		t.Error("String formatting error result was ", str, " expected ", expectedStr)
	}
}

func TestStep(t *testing.T) {
	l := NewLife(3, 3, stepInput, Next)
	l.Step()
	str := l.String()
	if strings.Compare(l.String(), expectedStepStr) != 0 {
		t.Error("String formatting error result was ", str, " expected ", expectedStepStr)
	}
}
