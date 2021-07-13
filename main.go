package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	var privVal float64

	for {
		z -= (z*z - x) / (2 * z)
		if z == privVal || math.Abs(privVal-z) < 1e-13 {
			break
		} else {
			privVal = z
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
