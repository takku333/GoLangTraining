// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import "fmt"

func main() {
	ss := []string{"ab", "ab", "c", "de", "de", "ab"}
	fmt.Println(ss)
	fmt.Println("remove overlap")
	fmt.Println(rmOverlap(ss))
}

func rmOverlap(ss []string) []string {
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			ss = remove(ss, i)
			i--
		}
	}
	return ss
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
