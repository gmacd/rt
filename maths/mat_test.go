package maths

import (
	"testing"
	"unsafe"

	"github.com/gmacd/rt/test"
)

func TestMatSizeCorrect(t *testing.T) {
	m := NewMatIdent()
	test.Equal(t, 64, int(unsafe.Sizeof(m)), "Mat wrong size")
}

func TestIdent(t *testing.T) {
	m1 := NewMatIdent()
	m2 := NewMat(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
	test.Equal(t, m1, m2, "Identity not correct")		
}