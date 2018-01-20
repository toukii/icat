package icat

import (
	"os"
	"testing"

	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
)

func TestIsvg(t *testing.T) {
	bs := goutils.ReadFile("test_images/github.svg")
	img, err := DecodeSVG(bytes.NewReader(bs))
	if err != nil {
		t.Errorf("%s", err)
	}
	ICat(img, os.Stdout)
}

func TestDisplay(t *testing.T) {
	if err := DisplaySVG(goutils.ReadFile("test_images/github.svg")); err != nil {
		t.Errorf("%s", err)
	}
}
