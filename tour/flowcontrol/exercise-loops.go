package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	delta := 1.0
	for delta > 1e-6 {
		pz := z
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(z - pz)
		fmt.Println(delta, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(102))
	fmt.Println(math.Sqrt(102))
}
