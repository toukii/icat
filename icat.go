package icat

import (
	"fmt"
	"image"
	"io"

	imgcat "github.com/martinlindhe/imgcat/lib"
	"github.com/oliamb/cutter"
)

func ICat(img image.Image, wr io.Writer) error {
	return imgcat.CatImage(img, wr)
}

func ICatRect(img image.Image, height, width int, wr io.Writer) error {
	bud := img.Bounds()
	fmt.Printf("img y:%d, x:%d\n", bud.Dy(), bud.Dx())
	if height <= 0 {
		height = bud.Dy()
	}
	if width <= 0 {
		width = bud.Dx()
	}
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  height,            // height in pixel or Y ratio(see Ratio Option below)
		Width:   width,             // width in pixel or X ratio
		Mode:    cutter.TopLeft,    // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{0, 0}, // Position of the top left point
		Options: 0,                 // Accepted Option: Ratio
	})
	if err != nil {
		return err
	}

	return ICat(cImg, wr)
}

// inFile := "file.jpg"
// // using a image.Image
// canvas := image.NewRGBA(image.Rect(0, 0, 20, 20))
// canvas.Set(10, 10, image.NewUniform(color.RGBA{255, 255, 255, 255}))
// imgcat.CatImage(canvas, os.Stdout)

// bd := img.Bounds()
// using a io.Reader
// f, _ := os.Open(inFile)
// imgcat.Cat(f, os.Stdout)

// // using filename
// imgcat.CatFile(inFile, os.Stdout)
