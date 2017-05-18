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
			img.Set(px, py, newton128(z))
		}
	}
	file128, err := os.Create("newton128.png")
	if err != nil {
		panic(err)
	}
	defer file128.Close()

	png.Encode(file128, img) // NOTE: ignoring errors

	img = image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton64(z))
		}
	}
	file64, err := os.Create("newton64.png")
	if err != nil {
		panic(err)
	}
	defer file64.Close()
	png.Encode(file64, img) // NOTE: ignoring errors

	// 	img = image.NewRGBA(image.Rect(0, 0, width, height))
	// 	for py := 0; py < height; py++ {
	// 		y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin)
	// 		for px := 0; px < width; px++ {
	// 			x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin)
	// 			// Image point (px, py) represents complex value z.
	// 			img.Set(px, py, newtonFloat(x, y))
	// 		}
	// 	}
	// 	fileFloat, err := os.Create("BigFloat.png")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer fileFloat.Close()
	// 	png.Encode(fileFloat, img) // NOTE: ignoring errors
}

func newton128(z complex128) color.Color {
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

func newton64(z complex64) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

// func newtonFloat(re *big.Float, im *big.Float) color.Color {
// 	const iterations = 37
// 	const contrast = 7
// 	for i := uint8(0); i < iterations; i++ {
// 		re2 := new(big.Float).Mul(re, re)
// 		re3 := new(big.Float).Mul(re2, re)
// 		im2 := new(big.Float).Mul(im, im)
// 		im3 := new(big.Float).Mul(im2, im)

// 		ri2 := new(big.Float).Mul(re, im2)
// 		r2i := new(big.Float).Mul(re2, im)

// 		threeRi2 := new(big.Float).Mul(ri2, big.NewFloat(3.0))
// 		threeR2i := new(big.Float).Mul(r2i, big.NewFloat(3.0))

// 		r3 := new(big.Float).Add(re3, (new(big.Float).Neg(threeRi2)))
// 		i3 := new(big.Float).Add(threeR2i, (new(big.Float).Neg(im3)))

// 		re = new(big.Float).Quo(new(big.Float).Add(re, new(big.Float).Neg(new(big.Float).Quo(big.NewFloat(1.0), r3))), big.NewFloat(4.0))
// 		im = new(big.Float).Quo(new(big.Float).Add(im, new(big.Float).Neg(new(big.Float).Quo(big.NewFloat(1.0), i3))), big.NewFloat(4.0))

// 		re4 := new(big.Float).Mul(new(big.Float).Mul(re, re), new(big.Float).Mul(re, re))
// 		im4 := new(big.Float).Mul(new(big.Float).Mul(im, im), new(big.Float).Mul(im, im))
// 		r2i2 := new(big.Float).Mul(new(big.Float).Mul(re, re), new(big.Float).Mul(im, im))
// 		r3i := new(big.Float).Mul(new(big.Float).Mul(re, re), new(big.Float).Mul(re, im))
// 		ri3 := new(big.Float).Mul(new(big.Float).Mul(re, im), new(big.Float).Mul(im, im))

// 		r4 := new(big.Float).Add(new(big.Float).Add(re4, new(big.Float).Neg(new(big.Float).Mul(r2i2, big.NewFloat(6.0)))), im4)
// 		i4 := new(big.Float).Add(new(big.Float).Mul(r3i, big.NewFloat(4.0)), new(big.Float).Neg(new(big.Float).Mul(ri3, big.NewFloat(4.0))))

// 		r4dif1 := new(big.Float).Add(r4, new(big.Float).Neg(big.NewFloat(1.0)))
// 		z4abs := new(big.Float)

// 		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
// 			return color.Gray{255 - contrast*i}
// 		}
// 	}
// 	return color.Black
// }
