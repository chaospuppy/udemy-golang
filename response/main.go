package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type myWriter struct{}

func (myWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return len(p), nil
}

func main() {
	resp, err := http.Get("https://google.com")
	var w myWriter
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// bs := make([]byte, 1)
	// resp.Body.Read(bs) // resp is of type io.ReadCloser, which is an interface to the Reader and Closer interfaces.  The Reader interface contains a function called `Read([]byte) (int, error)`
	//
	// fmt.Println(string(bs))
	io.Copy(w, resp.Body)
}
