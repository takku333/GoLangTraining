// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "set 2 args.\n")
		os.Exit(1)
	}
	if checkAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s : %s is anagram\n", os.Args[1], os.Args[1])
	} else {
		fmt.Printf("%s : %s is not anagram\n", os.Args[1], os.Args[1])
	}
}

func checkAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, c := range s1 {
		if !strings.Contains(s2, string(c)) {
			return false
		}
	}
	return true
}
