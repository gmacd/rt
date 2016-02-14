package maths

type Pos3 struct {
	X, Y, Z, _ float32
}

func NewPos3(x, y, z float32) Pos3 { return Pos3{X: x, Y: y, Z: z} }

func (p1 Pos3) Add(p2 Pos3) Vec3 {
	return NewVec3(p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z)
}
func (p1 Pos3) Sub(p2 Pos3) Vec3 {
	return NewVec3(p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z)
}

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

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z)
}
func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return NewVec3(v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z)
}

func (v Vec3) Addf(f float32) Vec3 { return NewVec3(v.X+f, v.Y+f, v.Z+f) }
func (v Vec3) Subf(f float32) Vec3 { return NewVec3(v.X-f, v.Y-f, v.Z-f) }
func (v Vec3) Mulf(f float32) Vec3 { return NewVec3(v.X*f, v.Y*f, v.Z*f) }
func (v Vec3) Divf(f float32) Vec3 { return v.Mulf(1.0 / f) }

func (v Vec3) Len() float32 { return Sqrtf(v.X*v.X + v.Y*v.Y + v.Z*v.Z) }
func (v Vec3) Norm() Vec3   { return v.Mulf(1.0 / v.Len()) }

func Dot(v1, v2 Vec3) float32 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func Cross(v1, v2 Vec3) Vec3 {
	return NewVec3(
		(v1.Y * v2.Z) - (v1.Z * v2.Y),
		(v1.Z * v2.X) - (v1.X * v2.Z),
		(v1.X * v2.Y) - (v1.Y * v2.X))
}

type Rgba struct {
	R, G, B, A float32
}

func NewRgba(r, g, b, a float32) Rgba { return Rgba{r, g, b, a} }

type Ray struct {
	Origin Pos3
	Dir    Vec3
}

func NewRay(o Pos3, dir Vec3) Ray { return Ray{o, dir} }
