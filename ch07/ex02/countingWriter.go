// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser

package main

import (
	"fmt"
	"io"
)

//CountWriter has writer and number of writing bytes
type CountWriter struct {
	w io.Writer
	n int64
}

func (cw *CountWriter) Write(p []byte) (n int, err error) {
	n, err = cw.w.Write(p)
	cw.n = int64(n)
	return n, err
}

//CountingWriter retrun new io.Writer and number of bytes written by Write()
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{w, 0}
	return &cw, &cw.n
}

func main() {
	var bc ByteCounter
	cbc, n := CountingWriter(&bc)
	fmt.Fprintf(cbc, "abc")
	fmt.Printf("n: %d\n", *n)
	fmt.Fprintf(cbc, "abcdef")
	fmt.Printf("n: %d\n", *n)

}
