// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
)

func main() {
	var arr [64]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)
	reverse(&arr)
	fmt.Println("reverse")
	fmt.Println(arr)
}

func reverse(s *[64]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
