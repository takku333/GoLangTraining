// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "a  あ \n阿\f\tA Σ  π"
	b := []byte(s)
	fmt.Println(s)
	b = encodeUTF8Space(b)
	fmt.Println(string(b))
}

func encodeUTF8Space(bs []byte) []byte {
	for i := 0; i < len(bs)-1; {
		r, size := utf8.DecodeRune(bs[i:])
		if unicode.IsSpace(r) {
			rNext, sizeNext := utf8.DecodeRune(bs[i+size:])
			if unicode.IsSpace(rNext) {
				bs = removeUTF8(bs, i+size, sizeNext)
				bs[i] = 0x20
				for j := 1; j < size; j++ {
					bs = remove(bs, i+j)
				}
				i++
				continue
			}
		}
		i += size
	}
	return bs
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeUTF8(slice []byte, i int, size int) []byte {
	for j := 0; j < size; j++ {
		copy(slice[i:], slice[i+1:])
	}
	return slice[:len(slice)-size]
}
