package icat

import (
	"fmt"
	"image"
	// "image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/oliamb/cutter"
)

func ICatImage(img image.Image) image.Image {
	w := NewEncodeWr(os.Stdout, nil)
	err := png.Encode(w, img)
	if err != nil {
		panic(err)
	}
	w.FlushStdout()
	return nil
}

func CatRectangle(img image.Image, minX, minY, maxX, maxY int, wr io.Writer) error {
	return CatRect(img, maxY-minY, maxX-minX, minY, minX, wr)
}

func CatRect(img image.Image, height, width, top, left int, wr io.Writer) error {
	bud := img.Bounds()
	// fmt.Printf("img y:%d, x:%d\n", bud.Dy(), bud.Dx())
	if height <= 0 {
		height = bud.Dy()
	}
	if width <= 0 {
		width = bud.Dx()
	}
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  height,                 // height in pixel or Y ratio(see Ratio Option below)
		Width:   width,                  // width in pixel or X ratio
		Mode:    cutter.TopLeft,         // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{left, top}, // Position of the top left point
		Options: 0,                      // Accepted Option: Ratio
	})
	if err != nil {
		return err
	}

	return Cat(cImg, wr)
}

func Cat(img image.Image, wr io.Writer) error {
	if typ, ok := wr.(*EncodeWr); ok {
		// encodErr := jpeg.Encode(typ, img, &jpeg.Options{100})
		encodErr := png.Encode(typ, img)
		flushErr := typ.FlushStdout()
		if flushErr != nil || encodErr != nil {
			return fmt.Errorf("err: %+v,%+v", flushErr, encodErr)
		}
		return nil
	}
	return png.Encode(wr, img)
}

func ICat(img image.Image, wr io.Writer) error {
	if typ, ok := wr.(*EncodeWr); ok {
		encodErr := png.Encode(typ, img)
		flushErr := typ.FlushStdout()
		if flushErr != nil || encodErr != nil {
			return fmt.Errorf("err: %+v,%+v", flushErr, encodErr)
		}
		return nil
	}
	w := NewEncodeWr(wr, nil)
	err := png.Encode(w, img)
	if err != nil {
		return err
	}
	if _, ok := wr.(*os.File); ok {
		return w.FlushStdout()
	}
	return w.Flush()
}

func ICatBase64(imgBase64 string, wr io.Writer) error {
	if typ, ok := wr.(*EncodeWr); ok {
		return typ.FlushBase64Stdout(imgBase64)
	}
	ew := NewEncodeWr(wr, nil)
	return ew.FlushBase64Stdout(imgBase64)
}

func ICatRect(img image.Image, height, width int, wr io.Writer) error {
	bud := img.Bounds()
	// fmt.Printf("img y:%d, x:%d\n", bud.Dy(), bud.Dx())
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

func ICatRead(r io.Reader, w io.Writer) error {
	if typ, ok := w.(*EncodeWr); ok {
		_, err := io.Copy(typ, r)
		if err != nil {
			return err
		}
		return typ.FlushStdout()
	}
	ew := NewEncodeWr(w, nil)
	_, err := io.Copy(ew, r)
	if err != nil {
		return err
	}
	return ew.FlushStdout()
}

func ICatHttp(uri string, w io.Writer) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	return ICatRead(resp.Body, os.Stdout)
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
