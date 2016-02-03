package scene

import (
	. "github.com/gmacd/rt/maths"
	"math"
)

type Camera struct {
	Pos Pos3
	Dir Vec3
	FovDeg float32
	cameraToWorld Mat
}

func NewCamera(pos Pos3, dir Vec3, fovDeg float32) *Camera {
	return &Camera{pos, dir, fovDeg, NewMatIdent()}
}

func (c *Camera) GenerateRay(x, y, w, h float32) Ray {
	// TODO Cache for each new frame
	var aspectRatio float32
	if (w >= h) {
		aspectRatio = w / h
	} else {
		aspectRatio = h / w
	}

	// Work out pixel positions in camera space
	fovFactor := float32(math.Tan(float64(DegToRad(2.0 * c.FovDeg))))
	camPixelPos := NewPos3(
		((2.0 * ((x + 0.5) / w)) - 1) * fovFactor * aspectRatio,
		(1 - (2.0 * ((y + 0.5) / h))) * fovFactor,
		-1)

	// Work out ray in world space, going from origin through camPix
	rayOriginWorld := c.cameraToWorld.MulPos3(Pos3{})
	rayDestWorld := c.cameraToWorld.MulPos3(camPixelPos)

	rayDirWorld := NewVecFrom2Pos3(rayOriginWorld, rayDestWorld)

	return NewRay(rayOriginWorld, rayDirWorld)
}
