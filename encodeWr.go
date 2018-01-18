package icat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

type EncodeWr struct {
	W   io.Writer
	buf *bytes.Buffer
	enc io.WriteCloser
}

func NewEncodeWr(w io.Writer, buff []byte) *EncodeWr {
	if buff == nil {
		buff = make([]byte, 0, 10240)
	}
	return &EncodeWr{
		W:   w,
		buf: bytes.NewBuffer(buff),
		enc: base64.NewEncoder(base64.StdEncoding, w),
	}
}

func (p *EncodeWr) Write(buf []byte) (n int, err error) {
	return p.buf.Write(buf)
}

func (p *EncodeWr) FlushStdout() error {
	fmt.Fprint(p.W, "\033]1337;File=;inline=1:")

	_, err := io.Copy(p.enc, p.buf)
	p.enc.Close()

	fmt.Fprintln(p.W, "\a")

	p.buf.Reset()
	return err
}

func (p *EncodeWr) FlushBase64Stdout(imgBase64 string) error {
	fmt.Fprint(p.W, "\033]1337;File=;inline=1:")
	defer fmt.Fprintln(p.W, "\a")
	_, err := fmt.Fprintln(p.W, imgBase64)
	return err
}

func (p *EncodeWr) Flush() error {
	_, err := io.Copy(p.W, p.buf)
	p.buf.Reset()
	return err
}
