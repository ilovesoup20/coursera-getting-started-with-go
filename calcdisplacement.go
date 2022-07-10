package main

import (
	"fmt"
	"math"
)

func main() {
	a, v0, s0 := 0.0, 0.0, 0.0

	fmt.Println("Enter acceleration:")
	fmt.Scan(&a)

	fmt.Println("Enter initial velocity:")
	fmt.Scan(&v0)

	fmt.Println("Enter initial distance:")
	fmt.Scan(&s0)

	fmt.Println(a, v0, s0)
	displacementFn := GenDisplaceFn(a, v0, s0)

	fmt.Println(displacementFn(1))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}
