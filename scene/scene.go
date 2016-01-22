package scene

import (
	. "github.com/gmacd/rt/maths"
)

type Sphere struct {
	Pos    Pos3
	Radius float32
}

func NewSphere(p Pos3, r float32) *Sphere {
	return &Sphere{p, r}
}

type Scene struct {
	camera  *Camera
	spheres []*Sphere
}

func NewScene() *Scene {
	return &Scene{
		nil,
		make([]*Sphere, 0)}
}

func (s *Scene) Camera() *Camera { return s.camera }

func (s *Scene) AddCamera(camera *Camera) {
	s.camera = camera
}

func (s *Scene) AddSpheres(sphere ...*Sphere) {
	s.spheres = append(s.spheres, sphere...)
}
