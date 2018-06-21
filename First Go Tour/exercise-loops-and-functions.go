package main

import (
	"fmt"
	"math"
)

const E = 0.0001

func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := 0.0
	
	for z - lastZ > E || z - lastZ < -E {
		lastZ = z
		z -= (z*z - x) / (2*z)
	}
	return z
	
}

func main() {
	for i:= 0.0; i < 11 ; i++ {
		fmt.Printf("Sqrt(%d) == %g\n", int(i), Sqrt(i))
		fmt.Printf("math.Sqrt(%d) - Sqrt(%d) == %g\n\n", int(i), int(i), math.Sqrt(i) - Sqrt(i))
	}
}

