package main

import (
	"fmt"

	. "github.com/gmacd/rt/maths"
	"github.com/gmacd/rt/scene"
	"github.com/gmacd/rt/support"
	"github.com/gmacd/rt/tracer"
)

func main() {
	fmt.Println("=== rt ===")

	renderer := support.NewGlRenderer(200, 200)
	renderer.Start()
	
	s := scene.CreateCornellBoxOfSpheres()

	rayTracer := tracer.NewRayTracer()

	for frame := range renderer.NextFrameChan() {
		if frame.ShouldStop {
			break
		}

		rayTracer.Render(s, frame)
		renderer.Render(frame)
	}
}
