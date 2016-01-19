package support

type imageRenderer struct {

}

func NewImageRenderer(filename string) {
	img, _ := os.Create(filename)
    defer img.Close()
    
    switch (
    jpeg.Encode(img, m, &jpeg.Options{jpeg.DefaultQuality})
}