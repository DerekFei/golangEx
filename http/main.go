package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp) //this is not the actual html in the resp

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs)) //this is the actual html in the resp

	//io.Copy(os.Stdout, resp.Body)  first arg has writer interface, second arg has reader interface

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
