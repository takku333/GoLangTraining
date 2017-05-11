// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "a  あ \n阿\tA Σ  π"
	fmt.Println(s)
	fmt.Println("reverse")
	bs := []byte(s)
	reverseUTF8(bs)
	fmt.Println(string(bs))
}

func reverseUTF8(bs []byte) {
	for i := 0; i < len(bs); {
		_, size := utf8.DecodeRune(bs[i:])
		if size > 1 {
			reverse(bs[i : i+size])
		}
		i += size
	}
	reverse(bs)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
