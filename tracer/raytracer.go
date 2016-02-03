package tracer

import (
	"math/rand"

	. "github.com/gmacd/rt/maths"
	"github.com/gmacd/rt/scene"
	"github.com/gmacd/rt/support"
)

type RayTracer struct {
}

func NewRayTracer() *RayTracer {
	return &RayTracer{}
}

func (rt *RayTracer) Render(scene *scene.Scene, frame *support.Frame) {
	pixels := frame.Pixels

	for i := 0; i < len(pixels); i++ {
		//r := NewRay(scene.Camera().Pos, 
	
		yellow := rand.Float32()/2.0 + 0.5
		pixels[i].R = yellow
		pixels[i].G = yellow
		pixels[i].B = 0
		pixels[i].A = 1.0
	}
}


/*type CameraRayGenerator struct {

}

func NewCameraRayTracer() CameraRayGenerator {
	return
}*/