package scene

import (
	. "github.com/gmacd/rt/maths"
)

func CreateCornellBoxOfSpheres() *Scene {
	s := NewScene()
	// Based on http://www.kevinbeason.com/smallpt/
	s.AddCamera(NewCamera(NewPos3(50, 52, 295.6), NewVec3(0, -0.042612, -1).Norm()))
	s.AddSpheres(
		NewSphere(NewPos3(0, 0, 0), 1e5), // Left
		NewSphere(NewPos3(0, 0, 0), 1e5), // Right
		NewSphere(NewPos3(0, 0, 0), 1e5), // Back
		NewSphere(NewPos3(0, 0, 0), 1e5), // Front
		NewSphere(NewPos3(0, 0, 0), 1e5), // Bottom
		NewSphere(NewPos3(0, 0, 0), 1e5)) // Top
	//s.AddLight(
	return s
}
