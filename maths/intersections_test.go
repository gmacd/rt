package maths

import "testing"

func TestRayTriangleIntersections(t *testing.T) {
	tri := Triangle{
		NewPos3(0, 1, 1),
		NewPos3(-1, -1, 1),
		NewPos3(1, -1, 1)}
		
	// Test hit - tri in front of ray, no maxT
	r := NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, _ := IntersectRayTriangle(&r, &tri, -1)
	if !hit {
		t.Errorf("Ray %v didn't intersect triangle %v", r, tri)
	}
	
	// Test hit - tri in front of ray, result less than maxT
	r = NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, _ = IntersectRayTriangle(&r, &tri, 2)
	if !hit {
		t.Errorf("Ray %v didn't intersect triangle %v", r, tri)
	}
	
	// Test miss - tri in front of ray, result greater than than maxT
	r = NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, _ = IntersectRayTriangle(&r, &tri, 0.5)
	if hit {
		t.Errorf("Ray %v intersected triangle %v", r, tri)
	}
	 
	// Test miss - ray in front of tri
	r = NewRay(NewPos3(0, 0, 2), NewVec3(0, 0, 1))
	hit, _ = IntersectRayTriangle(&r, &tri, -1)
	if hit {
		t.Errorf("Ray %v intersected triangle %v", r, tri)
	}
	
	// Test miss - tri in front of ray, but dir is directly up
	r = NewRay(NewPos3(0, 0, 0), NewVec3(1, 0, 0))
	hit, _ = IntersectRayTriangle(&r, &tri, -1)
	if hit {
		t.Errorf("Ray %v intersected triangle %v", r, tri)
	}
}
