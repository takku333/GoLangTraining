// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var bit = flag.Int("SHAbit", 256, "SHA digest bit select{256,384,512}")

func main() {
	flag.Parse()
	for _, s := range flag.Args() {
		switch *bit {
		case 256:
			fmt.Printf("%s [sha256] : %x\n", s, sha256.Sum256([]byte(s)))
		case 384:
			fmt.Printf("%s [sha348] : %x\n", s, sha512.Sum384([]byte(s)))
		case 512:
			fmt.Printf("%s [sha512] : %x\n", s, sha512.Sum512([]byte(s)))
		default:
			fmt.Printf("%s [sha256] : %x\n", s, sha256.Sum256([]byte(s)))
		}
	}
}
