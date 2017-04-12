// Author "Takumi Miyagawa"
// Copyright © 2017 Ricoh Co, Ltd. All rights reserved

package ex03

import (
	"testing"
	"bytes"
)


//BenchmarkEcho2 is echo2 benchmark test
func BenchmarkEcho2(b *testing.B) {
	buf := &bytes.Buffer{}
	ss := []string{"A","man,","a","plan,","a","canal:","Panama"}
	for i := 0; i < b.N; i++ {
		echo2(ss, buf)
	}
}


//BenchmarkEcho3 is echo3 benchmark test
func BenchmarkEcho3(b *testing.B) {
	buf := &bytes.Buffer{}
	ss := []string{"A","man,","a","plan,","a","canal:","Panama"}
	for i := 0; i < b.N; i++ {
		echo3(ss, buf)
	}
}