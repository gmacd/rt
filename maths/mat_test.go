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

func BenchmarkMatMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idx := float32(i)
		m1 := NewMat(
			idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7, idx+8,
			idx+9, idx+10, idx+11, idx+12, idx+13, idx+14, idx+15, idx+16)
		m2 := NewMat(
			idx+1000, idx+2000, idx+3000, idx+4000, idx+5000, idx+6000, idx+7000, idx+8000,
			idx+9000, idx+10000, idx+11000, idx+12000, idx+13000, idx+14000, idx+15000, idx+16000)
		
		m1.Mul(&m2)
	}
}
