package main

import (
	"fmt"
	"math/rand"

	"github.com/gmacd/rt/support"
)

func main() {
	fmt.Println("=== rt ===")

	renderer := support.NewGlRenderer(200, 200)
	renderer.Start()

	for frame := range renderer.NextFrameChan() {
		if frame.ShouldStop {
			break
		}

		pixels := frame.Pixels
		for i := 0; i < len(pixels); i += 4 {
			yellow := rand.Float32()/2.0 + 0.5
			pixels[i] = yellow
			pixels[i+1] = yellow
			pixels[i+2] = 0
			pixels[i+3] = 1.0
		}

		renderer.Render(frame)
	}
}
