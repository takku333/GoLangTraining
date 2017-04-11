// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"bufio"
)


func main() {
	fname := "compTimeReadHttp.txt"
	file, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file: %v\n", err)
		os.Exit(1)
	} 
	defer file.Close()
	writer := bufio.NewWriter(file)
	ch := make(chan string)
    for i := 1; i < 3; i++ {
		start := time.Now()
        writer.WriteString(fmt.Sprintf("loop%d\r\n", i))
        for _, url := range os.Args[1:] {
            go fetch(url, ch) // start a goroutine
        }
        for range os.Args[1:] {
            writer.WriteString(<-ch)
        }
    	writer.WriteString(fmt.Sprintf("%.2fs elapsed\r\n", time.Since(start).Seconds()))
    }
	writer.Flush()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\r\n", secs, nbytes, url)
}

//!-
