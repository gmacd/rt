package main

import (
	"fmt"

	"github.com/gmacd/rt/support"
	"github.com/gmacd/rt/tracer"
)

func main() {
	fmt.Println("=== rt ===")

	renderer := support.NewGlRenderer(200, 200)
	renderer.Start()

	rayTracer := tracer.NewRayTracer()

	for frame := range renderer.NextFrameChan() {
		if frame.ShouldStop {
			break
		}

		rayTracer.Render(frame)
		renderer.Render(frame)
	}
}
