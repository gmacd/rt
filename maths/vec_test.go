package maths

import (
	"testing"
	"unsafe"

	"github.com/gmacd/rt/test"
)

func TestSizesCorrect(t *testing.T) {
	// Allocate arrays so we ensure size in array is what's expected
	pos3Array := [...]Pos3{NewPos3(1, 2, 3), NewPos3(4, 5, 6)}
	expectedSize := len(pos3Array) * 16
	test.Equal(t, expectedSize, int(unsafe.Sizeof(pos3Array)), "Pos3 wrong size")

	vec3Array := [...]Vec3{NewVec3(1, 2, 3), NewVec3(4, 5, 6)}
	expectedSize = len(vec3Array) * 16
	test.Equal(t, expectedSize, int(unsafe.Sizeof(vec3Array)), "Vec3 wrong size")

	rgbaArray := [...]Rgba{NewRgba(1, 1, 1, 1), NewRgba(0.5, 0.5, 0.5, 1)}
	expectedSize = len(rgbaArray) * 16
	test.Equal(t, expectedSize, int(unsafe.Sizeof(rgbaArray)), "Rgba wrong size")
}

func TestDot(t *testing.T) {
	v1 := NewVec3(1, 2, 3)
	v2 := NewVec3(0, 0, 0)
	
	vr := Dot(v1, v2)
	if vr != 0 {
		t.Errorf("Expected 0, but found %v", vr)
	}
	
	v1 = NewVec3(1, 2, 3)
	v2 = NewVec3(4, 5, 6)
	
	vr = Dot(v1, v2)
	if vr != 32 {
		t.Errorf("Expected 32, but found %v", vr)
	}
}

func TestCross(t *testing.T) {
	// Test Parallel
	v1 := NewVec3(1, 0, 0)
	v2 := NewVec3(3, 0, 0)
	vr := Cross(v1, v2) 

	vrlen := vr.Len()
	if (vrlen != 0) {
		t.Errorf("Expected vrlen=%v, but found vrlen=%v", 0, vrlen)
	}
	
	// Test Perpendicular
	v1 = NewVec3(1, 0, 0)
	v2 = NewVec3(0, 1, 0)
	vr = Cross(v1, v2)
	
	expectedV := NewVec3(0, 0, 1)
	if (vr != expectedV) {
		t.Errorf("Expected %v, but found vr=%v", expectedV, vr)
	}
}
