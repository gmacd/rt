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