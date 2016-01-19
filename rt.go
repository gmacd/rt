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
		for i := 0; i < len(pixels); i++ {
			yellow := rand.Float32()/2.0 + 0.5
			pixels[i].R = yellow
			pixels[i].G = yellow
			pixels[i].B = 0
			pixels[i].A = 1.0
		}

		renderer.Render(frame)
	}
}
