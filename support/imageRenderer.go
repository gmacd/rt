package support

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageRenderer struct {
	Width, Height int
	Pixels        []float32
}

func NewImageRenderer(width, height int) *ImageRenderer {
	ir := &ImageRenderer{
		width, height,
		make([]float32, 4*width*height)}
	return ir
}

func (ir *ImageRenderer) RenderToFile(filename string) {
	img := image.NewNRGBA(image.Rect(0, 0, ir.Width, ir.Height))

	for y := 0; y < ir.Height; y++ {
		for x := 0; x < ir.Width; x++ {
			i := y*ir.Width*4 + x*4
			r := uint8(ir.Pixels[i] * 255.0)
			g := uint8(ir.Pixels[i+1] * 255.0)
			b := uint8(ir.Pixels[i+2] * 255.0)
			a := uint8(ir.Pixels[i+3] * 255.0)
			c := color.NRGBA{r, g, b, a}

			img.Set(x, y, c)
		}
	}

	imgWriter, _ := os.Create(filename)
	defer imgWriter.Close()

	var err error
	switch filepath.Ext(filename) {
	case ".png":
		err = png.Encode(imgWriter, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(imgWriter, img, &jpeg.Options{jpeg.DefaultQuality})
	}

	if err != nil {
		fmt.Printf("Error writing out %v: %v", filename, err)
	}
}
