package icat

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	img image.Image
)

func init() {
	fd, err := os.Open("test_images/gosea.jpg")
	if err != nil {
		panic(err)
	}
	img, err = jpeg.Decode(fd)
	if err != nil {
		panic(err)
	}
}

func gobike(rw http.ResponseWriter, req *http.Request) {
	ICat(img, rw)
}

func TestICat(t *testing.T) {
	for i := 0; i < 10; i++ {
		ICat(img, os.Stdout)
	}
	time.Sleep(2e9)
}

func TestICatBase64(t *testing.T) {
	ICatBase64("iVBORw0KGgoAAAANSUhEUgAAAAkAAAAJAQMAAADaX5RTAAAAA3NCSVQICAjb4U/gAAAABlBMVEX///+ZmZmOUEqyAAAAAnRSTlMA/1uRIrUAAAAJcEhZcwAACusAAArrAYKLDVoAAAAWdEVYdENyZWF0aW9uIFRpbWUAMDkvMjAvMTIGkKG+AAAAHHRFWHRTb2Z0d2FyZQBBZG9iZSBGaXJld29ya3MgQ1M26LyyjAAAAB1JREFUCJljONjA8LiBoZyBwY6BQQZMAtlAkYMNAF1fBs/zPvcnAAAAAElFTkSuQmCC", os.Stdout)
}

func TestICatHttp(t *testing.T) {
	ICatHttp("http://bramus.github.io/ws2-sws-course-materials/assets/xx/github.png", os.Stdout)
}

func TestEncodeWr(t *testing.T) {
	// fd, _ := os.OpenFile("out.txt", os.O_CREATE|os.O_RDWR, 0644)
	// ew := NewEncodeWr(fd, nil)
	ew := NewEncodeWr(os.Stdout, nil)
	for i := 0; i < 10; i++ {
		jpeg.Encode(ew, img, &jpeg.Options{100})
		if err := ew.FlushStdout(); err != nil {
			t.Errorf("%+v", err)
		}
	}
	time.Sleep(2e9)
}

func TestHttpRespImage(t *testing.T) {
	http.HandleFunc("/", gobike)
	http.ListenAndServe(":8080", nil)
}
