// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)
func main() {
	printDupFile(os.Stdout)
}

func printDupFile(w io.Writer) {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	count := 0
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}
	for line, fnames := range counts {
		if len(fnames) == 1 {
			for fname, n := range fnames {
				if n > 1 {
					fmt.Fprintf(w, "%s\n",line)
					fmt.Fprintf(w, "%s, %d\n", fname, n)
				}
			}		
		}else if len(fnames) > 1 {
			fmt.Fprintf(w,"%s\n",line)
			for fname, n := range fnames {				
				fmt.Fprintf(w, "%s, ", fname)
				count += n
			}
			fmt.Fprintf(w,"%d\n", count)
			count = 0
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
