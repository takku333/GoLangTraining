package main

import (
	"fmt"
	"time"
)

const max = 500000

func main() {
	for i := uint(1); i < max; i *= 2 {
		start, last := starter(i)
		s := time.Now()
		start <- true
		<-last
		e := time.Now().Sub(s)
		fmt.Printf("goroutin:%10d, time: %d\n", i, e.Nanoseconds())
	}
}

func starter(pNum uint) (chan bool, chan bool) {
	channels := make([]chan bool, pNum)
	lastCh := make(chan bool)

	for i := range channels {
		channels[i] = make(chan bool)
	}
	for i := range channels {
		if i < len(channels)-1 {
			go pipe(channels[i], channels[i+1])
		} else {
			go pipe(channels[i], lastCh)
		}
	}

	return channels[0], lastCh
}

func pipe(in <-chan bool, out chan<- bool) {
	ch := <-in
	out <- ch
}
