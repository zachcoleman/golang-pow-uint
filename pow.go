package main

import "golang.org/x/exp/constraints"

// PowUint provides a faster implementation of x^n where n is a uint
// the algorithm is O(log2(n))
func PowUint[T constraints.Integer | constraints.Float](x T, n uint) T {
	if x == 0 || n == 0 {
		return x
	}

	// start condition
	var zero T
	ret, curr := zero+1, x
	if n&1 == 1 {
		ret = ret * curr
	}
	n = n >> 1

	// start bit-shifting
	for n > 0 {
		curr = curr * curr
		if n&1 == 1 {
			ret = ret * curr
		}
		n = n >> 1
	}
	return ret
}
