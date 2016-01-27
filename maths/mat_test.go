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

func TestNewMatIdent(t *testing.T) {
	m1 := NewMatIdent()
	m2 := NewMat(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
	test.Equal(t, m1, m2, "Identity not correct")		
}

func TestNewMatScale(t *testing.T) {
	m1 := NewMatScale(2, 3, -4)
	m2 := NewMat(
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, -4, 0,
		0, 0, 0, 1)
	test.Equal(t, m1, m2, "Broken scale")
}

func TestMulIdent(t *testing.T) {
	m1 := NewMatIdent()
	m2 := NewMatIdent()
	m := m1.Mul(&m2)
	test.Equal(t, m, m1, "m is not identity")		
}

