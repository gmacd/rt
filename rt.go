package main

import (
	"fmt"

	"github.com/gmacd/rt/support"
)

const windowWidth = 600
const windowHeight = 600
const renderbufferWidth = 800
const renderbufferHeight = 600

func init() {
	// This is needed to arrange that main() runs on main thread.
	//runtime.LockOSThread()
}

func main() {
	fmt.Println("=== rt ===")

	renderer := support.NewGlRenderer(renderbufferWidth, renderbufferHeight)
	defer renderer.Shutdown()

	renderer.Init()
	renderer.Loop()
}
