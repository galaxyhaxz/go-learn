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
	const use = "Usage:\nmultiple [x] [y]\n"
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
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
	// doesn't seem to work?
	//paths := flag.Args()
	// if len(paths) < 3 {
		// flag.Usage()
		// os.Exit(1)
	// }

	veclib.TextFromAnotherLib()

	fmt.Printf("Calculating distance\n")

	acc := veclib.Vector2D_Acc(x, y)
	euc := veclib.Vector2D_Fast(x, y)
	fmt.Printf("The distance for (%d,%d) is\n", x, y)
	fmt.Printf("Squared: %d, Fast: %d\n", acc, euc)
	if acc < euc {
		acc ^= euc
		euc ^= acc
		acc ^= euc
	}
	diff := math.Abs(float64(acc - euc))
	diff /= float64(acc)
	diff *= 100
	fmt.Printf("Total error using fast algorithm: %.2f%%", diff)
}
