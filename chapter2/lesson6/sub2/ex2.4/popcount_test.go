package popcount

import "testing"

// pc[i] - кол-во единичный битов в i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	bench(b, PopCountShift)
}
