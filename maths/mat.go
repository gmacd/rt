package maths

type Mat struct {
	v[4][4]float32
}

func NewMat(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float32) Mat {
	return Mat{[4][4]float32 {
		{a, b, c, d},
		{e, f, g, h},
		{i, j, k, l},
		{m, n, o, p}}}
}

func NewMatIdent() Mat {
	return Mat{[4][4]float32 {
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0}}}
}

