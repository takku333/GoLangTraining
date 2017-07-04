// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//Reader implements io.Reader
type Reader string

func main() {
	url := "https://golang.org"
	_, err := html.Parse(NewReader(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %s, %v\n", url, err)
		os.Exit(1)
	}

}

//NewReader is return io.Reader using s
func NewReader(s string) *Reader {
	r := Reader(s)
	return &r
}

func (r *Reader) Read(p []byte) (int, error) {
	*r = Reader(p)
	return len(p), nil
}
