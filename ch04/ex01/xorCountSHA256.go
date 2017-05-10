// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "set 2 args.\n")
		os.Exit(1)
	}
	fmt.Println(xorCountSHA256(os.Args[1], os.Args[2]))
}

func xorCountSHA256(s1 string, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	n := 0
	for i, b1 := range c1 {
		n += PopCountLoop(uint64(b1) ^ uint64(c2[i]))
	}
	return n
}
