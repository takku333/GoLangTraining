// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser
package main

import "fmt"

func main() {
	res := increment(1)
	fmt.Printf("res: %d\n", res)
}

func increment(x int) (y int) {
	defer func() {
		if p := recover(); p == x {
			y = x + 1
		}
	}()
	panic(x)
}