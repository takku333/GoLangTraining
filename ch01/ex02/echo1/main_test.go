// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"testing"
	"bytes"
	"os"
)

func TestPrintCommandIndex(t *testing.T) {
	tests := []struct {
        args    []string
        want 	string
    }{
        {
            args: []string{"main.go","a","b"},
            want: "1 a\n2 b\n",
        },
    }
    for i, test := range tests {
		buf := &bytes.Buffer{}
		os.Args = test.args
        printCommandIndex(buf)
		result := buf.String()
        if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}