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
	if m1 != m2 {
		t.Errorf("Expected identity, but got: %v", m2)
	}
}

func TestNewMatScale(t *testing.T) {
	m1 := NewMatScale(2, 3, -4) 
	expected := NewMat(
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, -4, 0,
		0, 0, 0, 1)
	if m1 != expected {
		t.Errorf("Expected %v, but got: %v", expected, m1)
	}
}

func TestMulIdent(t *testing.T) {
	i := NewMatIdent()
	m := i.MulMat(&i)
	if m != i {
		t.Errorf("Expected %v, but got: %v", i, m)
	}
}

func TestMuls(t *testing.T) {
	i := NewMatIdent()
	mr := i.Mulf32(5)
	expected := NewMat(
		5, 0, 0, 0,
		0, 5, 0, 0,
		0, 0, 5, 0,
		0, 0, 0, 5)
	
	if mr != expected {
		t.Errorf("Expected %v, but found %v.", expected, mr)
	}
}

func TestMulv(t *testing.T) {
	m := NewMatIdent()
	v := NewVec3(10, 20, 30)

	vr := m.MulVec3(v)
	expected := NewVec3(10, 20, 30)
	if vr != expected {
		t.Errorf("Expected %v, but found %v.", expected, vr)
	}

	m = NewMatScale(2, 3, 4)
	vr = m.MulVec3(v)
	expected = NewVec3(20, 60, 120)
	if vr != expected {
		t.Errorf("Expected %v, but found %v.", expected, vr)
	}
}

// Benchmarks

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
		mr = m1.MulMat(&m2)
	}
	matResult = mr
}

func BenchmarkMatMuls(b *testing.B) {
	m1 := NewMat(
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16)
		
	b.ResetTimer()
	
	var mr Mat
	for i := 0; i < b.N; i++ {
		mr = m1.Mulf32(float32(i))
	}
	matResult = mr
}
