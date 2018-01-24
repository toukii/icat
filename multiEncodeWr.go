package icat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

type MultiEncodeWr struct {
	W   []io.Writer
	buf *bytes.Buffer
	enc io.WriteCloser
}

func NewMultiEncodeWr(buff []byte, ws ...io.Writer) *MultiEncodeWr {
	if buff == nil {
		buff = make([]byte, 0, 10240)
	}
	return &MultiEncodeWr{
		W:   ws,
		buf: bytes.NewBuffer(buff),
	}
}

func (p *MultiEncodeWr) Write(buf []byte) (n int, err error) {
	return p.buf.Write(buf)
}

func (p *MultiEncodeWr) FlushStdout(i int) error {
	if i >= len(p.W) {
		return fmt.Errorf("ws' len is:%d", len(p.W))
	}

	fmt.Fprint(p.W[i], "\033]1337;File=;inline=1:")

	enc := base64.NewEncoder(base64.StdEncoding, p.W[i])
	defer enc.Close()
	r := bytes.NewReader(p.buf.Bytes())
	_, err := io.Copy(enc, r)

	fmt.Fprintln(p.W[i], "\a")
	return err
}

func (p *MultiEncodeWr) FlushBase64Stdout(imgBase64 string) error {
	fmt.Fprint(p.W[0], "\033]1337;File=;inline=1:")
	defer fmt.Fprintln(p.W[0], "\a")
	_, err := fmt.Fprintln(p.W[0], imgBase64)
	return err
}

func (p *MultiEncodeWr) Close() {
	p.buf.Reset()
}

func (p *MultiEncodeWr) Flush(i int) error {
	if i >= len(p.W) {
		return fmt.Errorf("ws' len is:%d", len(p.W))
	}
	r := bytes.NewReader(p.buf.Bytes())
	_, err := io.Copy(p.W[i], r)
	return err
}
