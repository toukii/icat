package icat

import (
	"os"
	"testing"
	"time"

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

func TestGitHubSVG(t *testing.T) {
	un := time.Now().Unix()
	for i := 0; i < 10; i++ {
		err := GitHubSVG(un + int64(i))
		if err != nil {
			t.Error(err)
		}
	}
}
