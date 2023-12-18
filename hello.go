package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for z := 1.0; z < 1000000; z += 1 {
		if math.Pow(z, z) >= x {
			fmt.Println(math.Pow(z, z))
			fmt.Println("Found square root")
			fmt.Println(z)
			fmt.Println(z - 1)
			return z - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
}
