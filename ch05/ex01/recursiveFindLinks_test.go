// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"
	"testing"

	"reflect"

	"golang.org/x/net/html"
)

func TestRecursiveVisit(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "testdata/golang.html",
		},
	}
	for i, test := range tests {
		file, err := os.Open(test.input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s :%v\n", test.input, err)
			os.Exit(1)
		}
		defer file.Close()

		doc, err := html.Parse(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}
		test.want = visit(nil, doc)

		result := recursiveVisit(nil, doc)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
		}
	}

}
