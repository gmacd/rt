package maths

import (
	"math"
	"testing"
)

func TestRayTriangleIntersections(t *testing.T) {
	tri := Triangle{
		NewPos3(0, 1, 1),
		NewPos3(-1, -1, 1),
		NewPos3(1, -1, 1)}
		
	// Test hit - tri in front of ray, no maxT
	r := NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, tHit := IntersectRayTriangle(&r, &tri, math.MaxFloat32)
	if !hit {
		t.Errorf(
			"Tri in front of ray with no maxT. Should intersect. Ray: %v, Tri: %v, t: %v",
			r, tri, tHit)
	}
	
	// Test hit - tri in front of ray, result less than maxT
	r = NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, tHit = IntersectRayTriangle(&r, &tri, 2)
	if !hit {
		t.Errorf(
			"Tri in front of ray with maxT. Should intersect. Ray: %v, Tri: %v, t: %v",
			r, tri, tHit)
	}
	
	// Test miss - tri in front of ray, result greater than than maxT
	r = NewRay(NewPos3(0, 0, 0), NewVec3(0, 0, 1))
	hit, tHit = IntersectRayTriangle(&r, &tri, 0.5)
	if hit {
		t.Errorf(
			"Tri in front of ray with maxT. Shouldn't intersect. Ray: %v, Tri: %v, t: %v",
			r, tri, tHit)
	}
	 
	// Test miss - ray in front of tri
	r = NewRay(NewPos3(0, 0, 2), NewVec3(0, 0, 1))
	hit, tHit = IntersectRayTriangle(&r, &tri, math.MaxFloat32)
	if hit {
		t.Errorf(
			"Tri behind ray. Shouldn't intersect. Ray: %v, Tri: %v, t: %v",
			r, tri, tHit)
	}
	
	// Test miss - tri in front of ray, but dir is directly up
	r = NewRay(NewPos3(0, 0, 0), NewVec3(1, 0, 0))
	hit, tHit = IntersectRayTriangle(&r, &tri, math.MaxFloat32)
	if hit {
		t.Errorf(
			"Ray not pointing at tri. Shouldn't intersect. Ray: %v, Tri: %v, t: %v",
			r, tri, tHit)
	}
}
