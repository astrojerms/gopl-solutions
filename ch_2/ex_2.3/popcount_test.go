// Run with go test -bench=.
package popcount

import "testing"

// pc[i] is the population count i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the poplation count (number of set bits) of x.
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

func PopCountLoop(x uint64) int {
	var result int = 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// This extra function allows you to call the benchmarked functions with 0..N as argument values.
func caller(bm *testing.B, fn func(uint64) int) {
	for i := 0; i < bm.N; i++ {
		fn(uint64(i))
	}
}

// Benchmark functions should start with Benchmark so they are executed by the tester.
func BenchmarkPopCount(bm *testing.B) {
	caller(bm, PopCount)
}

func BenchmarkPopCountLoop(bm *testing.B) {
	caller(bm, PopCountLoop)
}
