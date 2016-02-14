package maths

type Triangle struct { 
	p1, p2, p3 Pos3
}

func IntersectRayTriangle(ray *Ray, tri *Triangle, maxT float32) (hit bool, t float32)  {
	// Vec from p1 to p2
	e1 := tri.p2.Sub(tri.p1)
	// Vec from p1 to p3
	e2 := tri.p3.Sub(tri.p1)
	s1 := Cross(ray.Dir, e2)
	divisor := Dot(s1, e1)
	if divisor == 0 {
		return false, 0
	}
	
	invDivisor := 1.0 / divisor
	
	// Calc barycentric coords
	d := ray.Origin.Sub(tri.p1)
	b1 := Dot(d, s1) * invDivisor
	if b1 < 0 || b1 > 1 {
		return false, 0
	}
	
	s2 := Cross(d, e1)
	b2 := Dot(d, s2) * invDivisor
	if b2 < 0 || b1 + b2 > 1 {
		return false, 0
	}
	
	// Line intersects triangle
	
	// Check that hit is less that maxT
	// TODO Add minT to Ray
	t = Dot(e2, s2) * invDivisor
	const rayMinT = 0.0
	if t < rayMinT || t > maxT {
		return false, 0
	}
	
	return true, t
}
