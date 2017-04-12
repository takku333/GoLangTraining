// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package main

import "fmt"

type meter float64
type feet float64
type kgram float64
type pound float64

func (mt meter) String() string { return fmt.Sprintf("%gm", mt) }
func (ft feet) String() string  { return fmt.Sprintf("%gft", ft) }
func (kg kgram) String() string { return fmt.Sprintf("%gkg", kg) }
func (pd pound) String() string { return fmt.Sprintf("%glb", pd) }


