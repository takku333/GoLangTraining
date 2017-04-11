// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep, indent := "", " ", "\n"
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		s += strconv.Itoa(i) + sep + arg + indent
	}
	fmt.Println(s)
}

//!-
