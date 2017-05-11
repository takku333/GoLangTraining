// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	types := make(map[string]int)
	invalid := 0 // count of invalid UTF-8 characters
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("open file: %v\n", err)
	}
	defer fp.Close()

	in := bufio.NewReader(fp)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsDigit(r) {
			types["Digit"]++
		} else if unicode.IsLetter(r) {
			types["Lettter"]++
		} else if unicode.IsSpace(r) {
			types["Space"]++
		}
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\type\tcount\n")
	for s, n := range types {
		fmt.Printf("%s\t%d\n", s, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
