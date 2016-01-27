package maths

import "math"

func Sqrtf(f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

func DegToRad(d float32) float32 { return d * math.Pi / 180.0 }
func RadToDeg(r float32) float32 { return r * 180.0 / math.Pi }
