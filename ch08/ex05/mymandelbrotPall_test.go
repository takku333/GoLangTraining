package main

import "testing"

func BenchmarkMymandelbrotPall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymandelbrotPall()
	}
}

func BenchmarkMymandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymandelbrot()
	}
}
