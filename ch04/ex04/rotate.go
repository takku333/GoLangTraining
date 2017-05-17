// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr = rotate(arr, 2)
	fmt.Println("rotete")
	fmt.Println(arr)
}

func rotate(arr []int, n int) []int {
	arr = append(arr[n:], arr[:n]...)
	return arr
}
