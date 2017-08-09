package popcountbench

import "testing"

func benchmark(f func(uint64) int, num uint64, loop int) {
	for i := 0; i < loop; i++ {
		f(num)
	}
}

func benchmarkPopcount(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		benchmark(PopCount, num, loop)
	}
}

func benchmarkPopCountBitClear(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		benchmark(PopCountBitShift, num, loop)
	}
}

func benchmarkPopcountBitClerar(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		benchmark(PopCountBitClear, num, loop)
	}
}

/*F1*/
func BenchmarkPopcountF1L10(b *testing.B) {
	benchmarkPopcount(b, 0xF, 10)
}

func BenchmarkPopCountBitClearF1L10(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xF, 10)
}

func BenchmarkPopcountBitClerarF1L10(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xF, 10)
}

func BenchmarkPopcountF1L10000(b *testing.B) {
	benchmarkPopcount(b, 0xF, 10000)
}

func BenchmarkPopCountBitClearF1L10000(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xF, 10000)
}

func BenchmarkPopcountBitClerarF1L10000(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xF, 10000)
}

/*F4*/
func BenchmarkPopcountF4L10(b *testing.B) {
	benchmarkPopcount(b, 0xFFFF, 10)
}

func BenchmarkPopCountBitClearF4L10(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xFFFF, 10)
}

func BenchmarkPopcountBitClerarF4L10(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xFFFF, 10)
}

func BenchmarkPopcountF4L10000(b *testing.B) {
	benchmarkPopcount(b, 0xFFFF, 10000)
}

func BenchmarkPopCountBitClearF4L10000(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xFFFF, 10000)
}

func BenchmarkPopcountBitClerarF4L10000(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xFFFF, 10000)
}

/*F8*/
func BenchmarkPopcountF8L10(b *testing.B) {
	benchmarkPopcount(b, 0xFFFFFFFF, 10)
}

func BenchmarkPopCountBitClearF8L10(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xFFFFFFFF, 10)
}

func BenchmarkPopcountBitClerarF8L10(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xFFFFFFFF, 10)
}

func BenchmarkPopcountF8L10000(b *testing.B) {
	benchmarkPopcount(b, 0xFFFFFFFF, 10000)
}

func BenchmarkPopCountBitClearF8L10000(b *testing.B) {
	benchmarkPopCountBitClear(b, 0xFFFFFFFF, 10000)
}

func BenchmarkPopcountBitClerarF8L10000(b *testing.B) {
	benchmarkPopcountBitClerar(b, 0xFFFFFFFF, 10000)
}
