// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var width, height int // canvas size in pixels
var cells int         // number of grid cells
var xyrange float64   // axis ranges (-xyrange..+xyrange)
var xyscale float64   // pixels per x or y unit
var zscale float64    // pixels per z unit
var angle float64     // angle of x, y axes (=30°)
var highcolor int64
var lowcolor int64
var sin, cos float64

func init() {
	width, height = 600, 320                    // canvas size in pixels
	cells = 100                                 // number of grid cells
	xyrange = 30.0                              // axis ranges (-xyrange..+xyrange)
	xyscale = float64(width) / 2 / xyrange      // pixels per x or y unit
	zscale = float64(height) * 0.4              // pixels per z unit
	angle = math.Pi / 8                         // angle of x, y axes (=30°)
	highcolor = 0xFF0000                        // color of High max
	lowcolor = 0x0000FF                         // color of Low min
	sin, cos = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
}

type fSurface func(float64, float64) float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, r)
}

func surface(out io.Writer, r *http.Request) {
	query := r.URL.Query()

	if len(query) != 0 {
		for qname, qvalue := range query {
			var err error
			switch qname {
			case "width":
				width, err = strconv.Atoi(qvalue[0])
				xyscale = float64(width) / 3.0 / float64(xyrange) // pixels per x or y unit
			case "height":
				height, err = strconv.Atoi(qvalue[0])
				zscale = float64(height) * 0.2 // pixels per z unit
			case "angle":
				deg, aerr := strconv.ParseInt(qvalue[0], 10, 0)
				angle = float64(deg) / 180 * math.Pi
				sin, cos = math.Sin(angle), math.Cos(angle)
				err = aerr
			case "highcolor":
				highcolor, err = strconv.ParseInt(qvalue[0], 16, 0)
			case "lowcolor":
				lowcolor, err = strconv.ParseInt(qvalue[0], 16, 0)
			default:
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "surface_server: %v\n", err)
			}
		}
	}
	outputSVG(out, fSinr)
}

func outputSVG(out io.Writer, f fSurface) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, isNaNa := corner(i+1, j, f)
			bx, by, bz, isNaNb := corner(i, j, f)
			cx, cy, cz, isNaNc := corner(i, j+1, f)
			dx, dy, dz, isNaNd := corner(i+1, j+1, f)
			if isNaNa || isNaNb || isNaNc || isNaNd {
				continue
			}
			aveZ := ((az + bz + cz + dz) / 4)
			color := calcColor(highcolor, lowcolor, aveZ)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%06X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f fSurface) (sx float64, sy float64, z float64, isNaN bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z = f(x, y)
	if math.IsNaN(z) {
		isNaN = true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = float64(width/2) + (x-y)*cos*xyscale
	sy = float64(height/2) + (x+y)*sin*xyscale - z*zscale

	return sx, sy, (z + 1) / 2, isNaN
}

func calcColor(highcolor int64, lowcoler int64, factor float64) int64 {
	r := int64(float64(getRFromRGB(highcolor)-getRFromRGB(lowcoler))*factor) + getRFromRGB(lowcoler)
	g := int64(float64(getGFromRGB(highcolor)-getGFromRGB(lowcoler))*factor) + getGFromRGB(lowcoler)
	b := int64(float64(getBFromRGB(highcolor)-getBFromRGB(lowcoler))*factor) + getBFromRGB(lowcoler)

	return r<<16 + g<<8 + b
}

func getRFromRGB(color int64) int64 {
	return (color & 0xff0000) >> 16
}

func getGFromRGB(color int64) int64 {
	return (color & 0x00ff00) >> 8
}

func getBFromRGB(color int64) int64 {
	return (color & 0x0000ff)
}
