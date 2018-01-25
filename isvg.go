package icat

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"text/template"

	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
	"github.com/toukii/icat/svg"
)

func GitHubSVG(src int64) error {
	rd := rand.New(rand.NewSource(src))

	data := make(map[string]string)
	du := rd.Intn(360)
	// du = 180
	// fmt.Println(du)

	data["CatAroundC"] = fmt.Sprintf("%s", svg.GetC2(du))
	data["OuterC"] = fmt.Sprintf("%s", svg.GetC2(du+80))
	data["CatC"] = fmt.Sprintf("%s", svg.GetC2(du-80))
	// c := svg.GetC2(du).String()
	// data["CatAroundC"] = fmt.Sprintf("%s", c[0])
	// data["OuterC"] = fmt.Sprintf("%s", c[1])
	// data["CatC"] = fmt.Sprintf("%s", c[2])

	// tpl := template.New("svg/github.tpl")
	tpl := template.New("github")
	tpl, err := tpl.Parse(goutils.ToString(svg.MustAsset("svg/github.tpl")))
	if err != nil {
		return err
	}
	w := bytes.NewWriter(nil)
	err = tpl.Execute(w, data)
	if err != nil {
		return err
	}

	// fmt.Printf("%s", w.Bytes())
	// goutils.WriteFile("a.svg", w.Bytes())

	return DisplaySVG(w.Bytes())
}

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
