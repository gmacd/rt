package maths

import "testing"

func TestMuls(t *testing.T) {
	i := NewMatIdent()
	m := &Mat{}
	Muls(5, i, m)
	expected := NewMat(
		5, 0, 0, 0,
		0, 5, 0, 0,
		0, 0, 5, 0,
		0, 0, 0, 5)
	
	if *m != *expected {
		t.Errorf("Expected %v, but found %v.", *expected, *m)
	}
}
