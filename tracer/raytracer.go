package tracer

import (
	"math"

	. "github.com/gmacd/rt/maths"
	"github.com/gmacd/rt/scene"
	"github.com/gmacd/rt/support"
)

type RayTracer struct {
	rayGen cameraRayGenerator
}

func NewRayTracer() *RayTracer {
	return &RayTracer {
		newCameraRayTracer() }
}

func (rt *RayTracer) Render(scene *scene.Scene, frame *support.Frame) {
	pixels := frame.Pixels
	w, h := frame.Width, frame.Height

	rt.rayGen.prepareNewFrame(scene.Camera(), w, h)
	
	tri := Triangle{
		NewPos3(0, 1, 1),
		NewPos3(-1, -1, 1),
		NewPos3(1, -1, 1)}

	pixelIdx := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := rt.rayGen.GeneratePrimaryRay(float32(x), float32(y))
			
			hit, _ := IntersectRayTriangle(&r, &tri, math.MaxFloat32)
			
			var yellow float32 = 0
			if hit {
				yellow = 1
			}
			pixels[pixelIdx].R = yellow
			pixels[pixelIdx].G = yellow
			pixels[pixelIdx].B = 0
			pixels[pixelIdx].A = 1.0

			pixelIdx++
		}
	}
}


// Generates primary rays.  Caching lots of data used when ray-casting
// for a single frame.
type cameraRayGenerator struct {
	cam *scene.Camera
	camToWorld Mat
	cameraOriginWorld Pos3

	aspectRatio float32
	fovFactor float32
	w, h float32
}

func newCameraRayTracer() cameraRayGenerator {
	return cameraRayGenerator {}
}

func (c *cameraRayGenerator) prepareNewFrame(
	cam *scene.Camera,
	w, h int) {

	c.cam = cam
	c.w = float32(w)
	c.h = float32(h)

	// TODO Build transform
	c.camToWorld = NewMatIdent()

	if (c.w >= c.h) {
		c.aspectRatio = c.w / c.h
	} else {
		c.aspectRatio = c.h / c.w
	}

	c.fovFactor = float32(math.Tan(float64(DegToRad(2.0 * cam.FovDeg))))

	c.cameraOriginWorld = c.camToWorld.MulPos3(Pos3{})
}

func (c *cameraRayGenerator) GeneratePrimaryRay(x, y float32) Ray {
	// Work out pixel positions in camera space
	camPixelPos := NewPos3(
		((2.0 * ((x + 0.5) / c.w)) - 1) * c.fovFactor * c.aspectRatio,
		(1 - (2.0 * ((y + 0.5) / c.h))) * c.fovFactor,
		-1)

	// Work out ray in world space, going from origin through camPix
	rayDestWorld := c.camToWorld.MulPos3(camPixelPos)
	rayDirWorld := NewVecFrom2Pos3(c.cameraOriginWorld, rayDestWorld)
	return NewRay(c.cameraOriginWorld, rayDirWorld)
}
