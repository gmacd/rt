package maths

type Triangle struct {
	p1, p2, p3 Pos3
}

func IntersectRayTriangle(ray *Ray, tri *Triangle, maxT float32) (hit bool, t float32)  {
	return false, 0
}
