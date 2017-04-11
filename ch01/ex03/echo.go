// Author "Takumi Miyagawa"
// Copyright © 2017 Ricoh Co, Ltd. All rights reserved

package ex03

import (
	"fmt"
	"os"
	"strings"
)

func echo2 (ss[] string) {
	s, sep := "", ""
	for _, arg := range ss[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3 (ss[] string) {
	fmt.Println(strings.Join(os.Args[1:], " "))
}