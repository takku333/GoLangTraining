// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 8         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type fSurface func(float64, float64) float64

func main() {
	for _, surface := range os.Args[1:] {
		switch surface {
		case "eggBox":
			outputSVG(fEggBox)
		case "mogle":
			outputSVG(fMogle)
		case "saddle":
			outputSVG(fSaddle)
		default:
			outputSVG(fSinr)
		}
	}
}

func outputSVG(f fSurface) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, isNaNa := corner(i+1, j, f)
			bx, by, isNaNb := corner(i, j, f)
			cx, cy, isNaNc := corner(i, j+1, f)
			dx, dy, isNaNd := corner(i+1, j+1, f)
			if isNaNa || isNaNb || isNaNc || isNaNd {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f fSurface) (sx float64, sy float64, isNaN bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) {
		isNaN = true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, isNaN
}

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
