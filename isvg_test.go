package icat

import (
	"os"
	"testing"

	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
)

func TestIsvg(t *testing.T) {
	bs := goutils.ReadFile("github.svg")
	dbs, err := Decode(bs)
	if err != nil {
		t.Errorf("%s", err)
	}
	if err := ICatRead(bytes.NewReader(dbs), os.Stdout); err != nil {
		t.Errorf("%s", err)
	}
}

func TestDisplay(t *testing.T) {
	if err := Display(goutils.ReadFile("github.svg")); err != nil {
		t.Errorf("%s", err)
	}
}
