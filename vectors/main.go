package main

// basic program to demonstrate importing from another library
// - using local libraries
// - reading inputs from the command line
// - performing math

import (
	"fmt"
	"math"
	"flag"
	"os"
	"github.com/galaxyhaxz/go-learn/veclib"
	//"../veclib"
)

func usage() {
	const use = "Usage:\nvectors [x] [y]\n"
	fmt.Fprintln(os.Stderr, use)
	flag.PrintDefaults()
}

// Get the % difference between two integers
func GetDistError(a, b int) float64 {
	// a will be the greater of the two
	if a < b {
		a ^= b
		b ^= a
		a ^= b
	}
	diff := math.Abs(float64(a - b))
	diff /= float64(a)
	diff *= 100

	return diff
}

func main() {
	var (
		x int
		y int
	)

	flag.IntVar(&x, "x", 0, "X value")
	flag.IntVar(&y, "y", 0, "Y value")
	flag.Usage = usage
	flag.Parse()

	// Can't divide by zero
	if x == 0 && y == 0 {
		flag.Usage()
		os.Exit(1)
	}

	veclib.TextFromAnotherLib()

	fmt.Printf("Calculating distance\n")

	acc := veclib.Vector2D_Acc(x, y)
	euc := veclib.Vector2D_Fast(x, y)
	fmt.Printf("The distance for (%d,%d) is\n", x, y)
	fmt.Printf("Squared: %d, Fast: %d\n", acc, euc)
	fmt.Printf("Total error using fast algorithm: %.2f%%", GetDistError(acc, euc))
}
