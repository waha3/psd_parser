package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"psd_parser/psd"
)

func main() {
	// fileName := os.Args[1]
	b, err := ioutil.ReadFile("./static/test2.psd")
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(b)

	file := &psd.File{
		Buf: reader,
	}

	header := &psd.Header{}
	resources := &psd.ImageResourcesSetion{}
	header.ReadHeader(file)
	resources.ReadImageResourcesSetion(file)
}
