package maths

import (
	"testing"
	"unsafe"

	"github.com/gmacd/rt/test"
)

func TestSizesCorrect(t *testing.T) {
	test.Equal(t, 16, int(unsafe.Sizeof(Pos3{1, 2, 3})), "Wrong size")
	test.Equal(t, 16, int(unsafe.Sizeof(Vec3{1, 2, 3})), "Wrong size")
	test.Equal(t, 16, int(unsafe.Sizeof(Rgba{1, 2, 3, 4})), "Wrong size")
}