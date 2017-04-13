// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"
	"strconv"
	"io"
)

func main() {
	printCommandIndex(os.Stdout)
}

func printCommandIndex(w io.Writer) {
	var s, sep, indent string
	sep = " "
	indent = "\n"
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + sep + os.Args[i] + indent
	}
	fmt.Fprintf(w, s)
}

//!-
