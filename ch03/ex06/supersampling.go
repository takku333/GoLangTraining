// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var highcolor color.RGBA
var lowcolor color.RGBA

func init() {
	highcolor.R = 0xff
	highcolor.G = 0xff
	lowcolor.B = 0xff
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, superSampling(img)) // NOTE: ignoring errors
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

func superSampling(source *image.RGBA) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, source.Rect.Dx()/2, source.Rect.Dy()/2))
	for py := 0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			img.Set(px, py, average(source, px*2, py*2))
		}
	}
	return img
}

func average(source image.Image, px int, py int) color.Color {
	var colors [4]color.Color
	colors[0] = source.At(px, py)
	colors[1] = source.At(px+1, py)
	colors[2] = source.At(px, py+1)
	colors[3] = source.At(px+1, py+1)
	var sr, sg, sb uint32
	for _, c := range colors {
		r, g, b, _ := c.RGBA()
		sr += r
		sg += g
		sb += b
	}
	return color.RGBA{
		uint8(sr / 4),
		uint8(sg / 4),
		uint8(sb / 4),
		255,
	}
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
