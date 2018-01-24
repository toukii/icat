package icat

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"reflect"
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
	fmt.Printf("%+v,%s\n", buf, buf)
	return p.buf.Write(buf)
}

func (p *EncodeWr) FlushStdout() error {
	fmt.Fprint(p.W, "\033]1337;File=;inline=1:")

	_, err := io.Copy(p.enc, p.buf)
	p.enc.Close()

	fmt.Fprintln(p.W, "\a")

	// p.buf.Reset()
	return err
}

func (p *EncodeWr) close() error {
	return p.enc.Close()
}

func (p *EncodeWr) FlushBase64Stdout(imgBase64 string) error {
	defer p.close()
	fmt.Fprint(p.W, "\033]1337;File=;inline=1:")
	defer fmt.Fprintln(p.W, "\a")
	_, err := fmt.Fprintln(p.W, imgBase64)
	return err
}

func (p *EncodeWr) Flush() error {
	defer p.close()
	_, err := io.Copy(p.W, p.buf)
	p.buf.Reset()
	return err
}

type EncodeStdout struct {
	enc    io.WriteCloser
	writed bool
}

var (
	EOFB = []byte{174, 66, 96, 130}
)

func NewEncodeStdout() *EncodeStdout {
	return &EncodeStdout{
		enc: base64.NewEncoder(base64.StdEncoding, os.Stdout),
	}
}

func (p *EncodeStdout) Write(buf []byte) (n int, err error) {
	if !p.writed {
		fmt.Fprint(os.Stdout, "\033]1337;File=;inline=1:")
		p.writed = true
	}
	n, err = p.enc.Write(buf)
	if n == 4 && reflect.DeepEqual(buf, EOFB) {
		fmt.Fprintln(os.Stdout, "\a")
	}
	return n, err
}
