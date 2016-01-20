package scene

type Scene struct {
	camera *Camera
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) AddCamera(camera *Camera) {
	s.camera = camera
}
