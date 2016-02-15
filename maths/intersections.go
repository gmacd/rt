package maths

type Triangle struct { 
	P1, P2, P3 Pos3
}

func IntersectRayTriangle(ray *Ray, tri *Triangle, maxT float32) (hit bool, t float32)  {
	// Vec from P1 to P2
	e1 := tri.P2.Sub(tri.P1)
	// Vec from P1 to P3
	e2 := tri.P3.Sub(tri.P1)
	s1 := Cross(ray.Dir, e2)
	divisor := Dot(s1, e1)
	if divisor == 0 {
		return false, 0
	}
	
	invDivisor := 1.0 / divisor
	
	// Calc barycentric coords
	d := ray.Origin.Sub(tri.P1)
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
	
	// Check that hit is greater than minT, but less that maxT
	// Less than zero would mean behind the ray.
	// TODO Add minT to Ray
	t = Dot(e2, s2) * invDivisor
	const rayMinT = 0.0
	if t < rayMinT || t > maxT {
		return false, 0
	}
	
	return true, t
}
