// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	printCommandName(os.Stdout)
}

func printCommandName(w io.Writer) {
	var s, sep string
	s += os.Args[0]
	sep = " "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Fprintf(w, s)
}

//!-
