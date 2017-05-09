// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

func commaFloat(s string) string {
	var buf bytes.Buffer
	var deci bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		s = s[1:]
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		deci.WriteString(s[dot:])
		s = s[:dot]
	}
	buf = *comma(s, &buf)
	buf.WriteString(deci.String())
	return buf.String()
}

func comma(s string, buf *bytes.Buffer) *bytes.Buffer {
	for i := 0; i < len(s); i++ {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf
}
