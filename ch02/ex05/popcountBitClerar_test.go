// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package pcbitclear

import (
	"testing"
	"gopl.io/ch2/popcount"
)

func TestPopCountBitClear(t *testing.T) {
	tests := []struct {
        x    uint64
        want int
    }{
        {
            x: 0x15,
            want: 3,
        },
    }
    for i, test := range tests {
        result := PopCountBitClear(test.x)
        if test.want != result {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.want)
        }
    }
}

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x0123456789abcde)
	}
}

func BenchmarkPopcountBitClerar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitClear(0x0123456789abcde)
	}
}