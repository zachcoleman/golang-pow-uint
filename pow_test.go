package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestIntPow(t *testing.T) {
	// ints
	assert.Equal(t, PowUint(0, 1), 0)
	assert.Equal(t, PowUint(1, 1), 1)
	assert.Equal(t, PowUint(2, 7), 128)
	assert.Equal(t, PowUint(3, 4), 81)
	assert.Equal(t, PowUint(4, 4), 256)
	assert.Equal(t, PowUint(5, 4), 625)
	assert.Equal(t, PowUint(2, 20), 1048576)

	// float
	assert.True(t, math.Abs(PowUint(2., 2)-4.0) < 1e-7)
	assert.True(t, math.Abs(PowUint(0.1, 3)-0.001) < 1e-7)
	assert.True(t, math.Abs(PowUint(1.2, 10)-6.1917364224) < 1e-7)
}

func LoopPow[T constraints.Integer | constraints.Float](x T, n uint) T {
	var zero T
	ret := zero + 1
	for i := 0; i < int(n); i++ {
		ret *= x
	}
	return ret
}

func BenchmarkPowUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 1; x < 4; x++ {
			for n := 0; n < 20; n++ {
				_ = PowUint(x, uint(n))
			}
		}
		for x := 1.; x < 4; x += 0.2 {
			for n := 0; n < 20; n++ {
				_ = PowUint(x, uint(n))
			}
		}
	}
}

func BenchmarkLoopPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 1; x < 4; x++ {
			for n := 0; n < 20; n++ {
				_ = LoopPow(x, uint(n))
			}
		}
		for x := 1.; x < 4; x += 0.2 {
			for n := 0; n < 20; n++ {
				_ = LoopPow(x, uint(n))
			}
		}
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 1; x < 4; x++ {
			for n := 0; n < 20; n++ {
				_ = int(math.Pow(float64(x), float64(n)))
			}
		}
		for x := 1.; x < 4; x += 0.2 {
			for n := 0; n < 20; n++ {
				_ = int(math.Pow(float64(x), float64(n)))
			}
		}
	}
}
