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
	test.Equal(t, m1, m, "m is not identity")		
}

var matResult Mat

func BenchmarkMatMul(b *testing.B) {
	m1 := NewMat(
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16)
	m2 := NewMat(
		1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
		9000, 10000, 11000, 12000, 13000, 14000, 15000, 16000)
		
	b.ResetTimer()
	
	var mr Mat
	for i := 0; i < b.N; i++ {
		mr = m1.Mul(&m2)
	}
	matResult = mr
}
