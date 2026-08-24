package primitives

import (
	"fmt"
)

func X(dst, target []uint64) {
	if len(dst) != 8 || len(target) != 8 {
		panic(fmt.Errorf("primitives.X: unexpected slice lengths: %d, %d expected: 8", len(dst), len(target)))
	}

	// The compiler can optimize this fixed-size loop where supported,
	// while CPUs without AVX2 can safely execute it as well.
	for i := range dst {
		dst[i] ^= target[i]
	}
}
