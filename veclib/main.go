package veclib

// library to implement vector and other math functions

import (
	"fmt"
	"math"
)

func TextFromAnotherLib() {
	fmt.Print("Initializing math from another dimension\n\n")
}

// Get an accurate length for 2D vector using square root
func Vector2D_Acc(x int, y int) int {
	return int(math.Sqrt(float64(x * x + y * y)))
}

// Approximate length of 2D vector using Euclidean distance
// Maximum error of 4.0%
// C version in "Graphics Gems IV", Academic Press, 1994
func Vector2D_Fast(x int, y int) int {
	var t int
	// always use positive integer
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	// make sure x is greatest
	if x < y {
		x ^= y
		y ^= x
		x ^= y
	}
	// crazy voodoo :P
	t = y + (y >> 1)
	return x - (x >> 5) - (x >> 7) + (t >> 2) + (t >> 6)
}
