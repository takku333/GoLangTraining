// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//RGB is red green blue 8bit coler code
type RGB struct {
	R uint8
	G uint8
	B uint8
}

var palette = []color.Color{ color.Black }
var colorStart = RGB{R:0x00, G:0x00, B:0xff}
var colorEnd = RGB{R:0xff, G:0x00, B:0x00}

func init(){
	for i:=0.0; i<1.0; i+=0.1 {
		R := uint8(float64(colorEnd.R) - float64(colorStart.R) * i) + colorStart.R
		G := uint8(float64(colorEnd.G) - float64(colorStart.G) * i) + colorStart.G
		B := uint8(float64(colorEnd.B) - float64(colorStart.B) * i) + colorStart.B
		palette = append(palette, color.RGBA{R, G, B, 0xff})
	} 	
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%(len(palette)-1))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
