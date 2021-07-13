package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d - %.15f\n", i, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println("My___Sqrt - ", Sqrt(2))
	fmt.Println("Math_Sqrt - ", math.Sqrt(2))
}
