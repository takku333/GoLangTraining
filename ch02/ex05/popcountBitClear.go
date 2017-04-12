// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package pcbitclear

// PopCountBitClear returns the population count (number of set bits) of x.
func PopCountBitClear(x uint64) int {
	n := 0
	for x != 0x0 {
		x = x & (x - 1)
		n++
	}
	return n
}
