// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

var highcolor color.RGBA
var lowcolor color.RGBA

func init() {
	highcolor.R = 0xff
	highcolor.G = 0xff
	lowcolor.B = 0xff
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	makeImage(w, r)
}

func makeImage(out io.Writer, r *http.Request) {
	var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
	width, height := 1024, 1024
	var dx, dy, zoom float64

	query := r.URL.Query()
	if len(query) != 0 {
		for qname, qvalue := range query {
			var err error
			switch qname {
			case "x":
				dx, err = strconv.ParseFloat(qvalue[0], 64)
			case "y":
				dy, err = strconv.ParseFloat(qvalue[0], 64)
			case "zoom":
				zoom, err = strconv.ParseFloat(qvalue[0], 64)
			default:
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "fractalServer: %v\n", err)
			}
		}
	}

	xmin -= zoom
	ymin -= zoom
	xmax += zoom
	ymax += zoom

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x+dx, y+dy)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				255 - contrast*n*(n%3),
				255 - contrast*n*(n%5),
				255 - contrast*n*(n%7),
				255,
			}
		}
	}
	return color.Black
}
