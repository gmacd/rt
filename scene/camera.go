package scene

import . "github.com/gmacd/rt/maths"

type Camera struct {
	Pos Pos3
	Dir Vec3
}

func NewCamera(pos Pos3, dir Vec3) *Camera {
	return &Camera{pos, dir}
}
