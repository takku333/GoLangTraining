// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package pcbitshift

// PopCountBitShift returns the population count (number of set bits) of x.
func PopCountBitShift(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 != 0 {
			n++
		}
	}
	return n
}
