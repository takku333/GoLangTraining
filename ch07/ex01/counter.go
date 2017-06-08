// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//ByteCounter is number of byte count
type ByteCounter int

//WordCounter is number of word count
type WordCounter int

//LineCounter is number of line count
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*wc++
	}
	return int(*wc), nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	for sc.Scan() {
		*lc++
	}
	return int(*lc), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	var wc WordCounter
	wc.Write([]byte("a ab abc abcde"))
	fmt.Println(wc) // "5", = len("hello")

	var lc LineCounter
	lc.Write([]byte("hello\nworld\n!!"))
	fmt.Println(lc) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

}
