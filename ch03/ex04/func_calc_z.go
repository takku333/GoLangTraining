// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import "math"

func fSinr(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func fEggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}

func fMogle(x, y float64) float64 {
	x2 := math.Pow(2, math.Sin(x))
	y2 := math.Pow(2, math.Sin(y))
	return (x2 * y2) / 20
}

func fSaddle(x, y float64) float64 {
	return (-math.Pow(x/xyrange, 2) + math.Pow(y/xyrange, 2))
}
