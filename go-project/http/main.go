package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	r, e := http.Get("http://google.com")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	lw := logWriter{}

	//bs := make([]byte, 99999)
	//r.Body.Read(bs)
	//fmt.Println(string(bs))
	io.Copy(lw, r.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes Written:" + string(len(bs)))
	return len(bs), nil
}
