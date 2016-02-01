package maths

import (
	"math"
)

func NewMat(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float32) *Mat {
	return &Mat{
		{a, b, c, d},
		{e, f, g, h},
		{i, j, k, l},
		{m, n, o, p}}
}

func NewMatZero() *Mat {
	return &Mat{}
}

func NewMatIdent() *Mat {
	return &Mat{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0}}
}

func NewMatScale(x, y, z float32) *Mat {
	return &Mat{
		{x, 0.0, 0.0, 0.0},
		{0.0, y, 0.0, 0.0},
		{0.0, 0.0, z, 0.0},
		{0.0, 0.0, 0.0, 1.0}}
}

// Left-handed (z into screen is +ve, y up is +ve, x right is +ve
func NewMatPerspective(fov, near, far float32) *Mat {
	// TODO Optimize to build mat directly
	i33 := far / (far - near)
	i34 := i33 * -near
	perspMat := &Mat{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, i33, i34},
		{0.0, 0.0, 1.0, 0.0}}

	invTanAngle := float32(1.0 / math.Tan(float64(DegToRad(fov)/2.0)))
	scaleMat := NewMatScale(invTanAngle, invTanAngle, 1)
	result := scaleMat.Mul(perspMat)
	return result
}
