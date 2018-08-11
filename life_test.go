package golife

import (
	"testing"
	"strings"
)

var (
	input    = []bool{false, false, false, true, true, true, false, true, false}
	expected = "0 0 0 1 1 1 0 1 0"
)

func TestString(t *testing.T) {
	l := NewLife(3, 3, input, Next)
	str := l.String()
	if strings.Compare(str, expected) != 0 {
		t.Error("String formatting error result was ", str, " expected ", expected)
	}
}
