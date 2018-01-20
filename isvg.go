package icat

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/toukii/bytes"
)

func Decode(bs []byte) ([]byte, error) {
	cmd := exec.Command("rsvg-convert")

	w := bytes.NewWriter(make([]byte, 0, 1024))
	cmd.Stdout = w
	cmd.Stdin = bytes.NewReader(bs)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

func Display(bs []byte) error {
	cmd := exec.Command("rsvg-convert")

	w := NewEncodeWr(os.Stdout, nil)
	defer w.FlushStdout()
	cmd.Stdout = w
	cmd.Stdin = bytes.NewReader(bs)
	return cmd.Run()
}

func ReadDisplay(r io.Reader) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return Display(bs)
}
