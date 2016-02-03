package maths

type Pos3 struct {
	X, Y, Z, _ float32
}

func NewPos3(x, y, z float32) Pos3 { return Pos3{X: x, Y: y, Z: z} }

type Vec3 struct {
	X, Y, Z, _ float32
}

func NewVec3(x, y, z float32) Vec3 { return Vec3{X: x, Y: y, Z: z} }

func NewVecFrom2Pos3(from, to Pos3) Vec3 {
	return Vec3 {
		X: to.X - from.X,
		Y: to.Y - from.Y,
		Z: to.Z - from.Z } 
}

func (v Vec3) Addf(f float32) Vec3 { return NewVec3(v.X+f, v.Y+f, v.Z+f) }
func (v Vec3) Subf(f float32) Vec3 { return NewVec3(v.X-f, v.Y-f, v.Z-f) }
func (v Vec3) Mulf(f float32) Vec3 { return NewVec3(v.X*f, v.Y*f, v.Z*f) }
func (v Vec3) Divf(f float32) Vec3 { return v.Mulf(1.0 / f) }

func (v Vec3) Len() float32 { return Sqrtf(v.X*v.X + v.Y*v.Y + v.Z*v.Z) }
func (v Vec3) Norm() Vec3   { return v.Mulf(1.0 / v.Len()) }

type Rgba struct {
	R, G, B, A float32
}

func NewRgba(r, g, b, a float32) Rgba { return Rgba{r, g, b, a} }

type Ray struct {
	Origin Pos3
	Dir    Vec3
}

func NewRay(o Pos3, dir Vec3) Ray { return Ray{o, dir} }
