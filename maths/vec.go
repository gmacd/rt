package maths

type Pos3 struct {
	X, Y, Z, _ float32
}

func NewPos3(x, y, z float32) Pos3 { return Pos3{X: x, Y: y, Z: z} }

type Vec3 struct {
	X, Y, Z, _ float32
}

func NewVec3(x, y, z float32) Vec3 { return Vec3{X: x, Y: y, Z: z} }

type Rgba struct {
	R, G, B, A float32
}

func NewRgba(r, g, b, a float32) Rgba { return Rgba{r, g, b, a} }
