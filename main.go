package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var privVal float64
	var i int = 1
	for {
		fmt.Printf("%d - %v\n", i, z)
		z -= (z*z - x) / (2 * z)
		if z == privVal || math.Abs(privVal-z) < 1e-13 {
			break
		} else {
			privVal = z
			i++
		}
	}
	return z
}

func main() {
	fmt.Println("My___Sqrt - ", Sqrt(2))
	fmt.Println("Math_Sqrt - ", math.Sqrt(2))
}
