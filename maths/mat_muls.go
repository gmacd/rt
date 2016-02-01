package maths

func Muls(s float32, m, mout *Mat)

// Temporarily here for benchmarking
func (m *Mat) MulsGo(s float32) Mat {
	return Mat{
		{ s*m[0][0], s*m[0][1], s*m[0][2], s*m[0][3] },
		{ s*m[1][0], s*m[1][1], s*m[1][2], s*m[1][3] },
		{ s*m[2][0], s*m[2][1], s*m[2][2], s*m[2][3] },
		{ s*m[3][0], s*m[3][1], s*m[3][2], s*m[3][3] }}
}
