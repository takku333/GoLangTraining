// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		log.Fatal("set text file")
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			counts := wordfreq(file)
			fmt.Println(file)
			for s, n := range counts {
				fmt.Printf("%q : %d\n", s, n)
			}
		}
	}
}

func wordfreq(file *os.File) map[string]int {
	counts := make(map[string]int)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		counts[sc.Text()]++
	}
	return counts
}
