package scene

import (
	"testing"

	. "github.com/gmacd/rt/maths"
)

func TestCreateSimpleScene(t *testing.T) {
	s := NewScene()
	// Based on http://www.kevinbeason.com/smallpt/
	s.AddCamera(NewCamera(NewPos3(50, 52, 295.6), NewVec3(0, -0.042612, -1).Norm()))
}
