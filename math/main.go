package main

import (
	"fmt"
	"math"
)

// go sample program
// demonstrate the use of:
//   - basic math
//   - switch case
//   - for loop
//   - checking variable type
//   - defining a function
//   - working with arrays

func PrintVarType(i interface{}) {
	switch t := i.(type) {
	case float64:
		fmt.Println("Type is float64")
	case bool:
		fmt.Println("Type is bool")
	case int:
		fmt.Println("Type is int")
	default:
		fmt.Println("Unknown type %T", t)
	}
}

func HijackDungeon(i int) {
	var dung [40][40]int

	// initialize the dungeon
	for x := 0; x < 40; x++ {
		for y := 0; y < 40; y++ {
			dung[x][y] = i + x + y
		}
	}

	if i * 3 < 40 { // no bounds overflow
		fmt.Println("Dungeon (", i * 3, ",7 ) is", dung[i * 3][7])
	}

	// dirty math
	tiles := [5]int { 1 * i, 2 + i, 3 << i, 4 ^ i, 500 / i }
	for j := 0; j < len(tiles); j++ {
		fmt.Printf("Tile %d is %d\n", j, tiles[j])
	}
}

func main() {
	// greet our user kindly!
	fmt.Printf("Welcome to the dark BDSM dungeon!\n")

	// initialize square root multiplier
	var xx [10]int
	for i := 0; i < len(xx); i++ {
		xx[i] = i
	}
	var f float64
	for i := 0; i < 10; i++ {
		if i == 4 || i == 7 {
			HijackDungeon(i)
		}
		var tmp int = i * xx[i]
		f = float64(tmp)
		if f != 0 {
			const equation string = "The square root of"
			fmt.Printf("%s %d is %f ", equation, tmp, math.Sqrt(f))
			switch i {
			case 3, 5, 7, 9:
				fmt.Printf("Odd iterator\n")
			case 2, 4, 6, 8:
				fmt.Printf("Even iterator\n")
			default:
				fmt.Printf("Other iterator\n")
			}
		}
	}

	// check what the variable is
	PrintVarType(f)

	// let's end this puppy!
	fmt.Println("Finished\n")
}
