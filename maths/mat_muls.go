package maths

func Muls(s float32, m, mout *Mat)

// Temporarily here for benchmarking
func MulsGo(s float32, m, mout *Mat) {
	mout[0][0] = s*m[0][0]
	mout[0][1] = s*m[0][1]
	mout[0][2] = s*m[0][2]
	mout[0][3] = s*m[0][3]
	mout[1][0] = s*m[1][0]
	mout[1][1] = s*m[1][1]
	mout[1][2] = s*m[1][2]
	mout[1][3] = s*m[1][3]	
	mout[2][0] = s*m[2][0]
	mout[2][1] = s*m[2][1]
	mout[2][2] = s*m[2][2]
	mout[2][3] = s*m[2][3]
	mout[3][0] = s*m[3][0]
	mout[3][1] = s*m[3][1]
	mout[3][2] = s*m[3][2]
	mout[3][3] = s*m[3][3]
}
