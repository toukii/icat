package icat

import (
	// "fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/toukii/bytes"
)

func DecodeSVG(r io.Reader) (image.Image, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("rsvg-convert")

	w := bytes.NewWriter(make([]byte, 0, 1024))
	cmd.Stdout = w
	cmd.Stdin = bytes.NewReader(bs)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return png.Decode(bytes.NewReader(w.Bytes()))
}

func DisplaySVG(bs []byte) error {
	cmd := exec.Command("rsvg-convert")

	// fmt.Printf("%s\n", bs)

	w := NewEncodeWr(os.Stdout, nil)
	defer w.FlushStdout()
	cmd.Stdout = w
	cmd.Stdin = bytes.NewReader(bs)
	return cmd.Run()
}

func ReadDisplaySVG(r io.Reader) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return DisplaySVG(bs)
}
