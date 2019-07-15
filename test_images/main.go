package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	// go http.Serve(ln, http.FileServer(FS))
	// go http.Serve(ln, http.FileServer(http.Dir("/Users/toukii/PATH/GOPATH/ezbuy/goflow/src/github.com/toukii/mdblog/MDFs")))
	go http.Serve(ln, http.FileServer(http.Dir(".")))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	<-make(chan bool)
}
