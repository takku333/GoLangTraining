// Author "Takumi Miyagawa"
// Copyright © 2017 Ricoh Co, Ltd. All rights reserved

package ex03

import (
	"fmt"
	"os"
	"strings"
	"io"
)

func echo2 (ss[] string, w io.Writer) {
	s, sep := "", ""
	for _, arg := range ss[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintf(w, s)
}

func echo3 (ss[] string, w io.Writer) {
	fmt.Fprintf(w, strings.Join(os.Args[1:], " "))
}