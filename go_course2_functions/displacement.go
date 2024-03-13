package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return s0 + (v0 * t) + (a * t * t / 2)
	}
}

func main() {
	var v0, s0, a, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scanf("%f", &v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scanf("%f", &s0)

	displaceFn := GenDisplaceFn(a, v0, s0)

	fmt.Print("Enter time: ")
	fmt.Scanf("%f", &t)

	fmt.Println(displaceFn(t))
}