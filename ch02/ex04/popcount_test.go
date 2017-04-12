// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package pcbitshift

import (
	"testing"

	"gopl.io/ch2/popcount"
)

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x0123456789abcde)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitShift(0x0123456789abcde)
	}
}
