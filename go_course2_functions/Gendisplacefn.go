package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)
func text(chains ...string) []float64 {
	var args []float64
	for _ , chain := range chains{
		fmt.Println("Please enter the " + chain)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		aux, _ := strconv.ParseFloat(scanner.Text(),64)
		args = append(args, aux)
	}
	return args
}
func GenDisplaceFn(a, v0,s0 float64) func(float64) float64{
	fn := func(t float64) float64{
		return (0.5*(a * (math.Pow(t,2)))) + (v0 * t) + s0
	}
	return fn
}
func main(){
	args := []string{"Aceleration", "Initial velocity", "Initial displacement"}
	argsFloats := text(args...)
	displacement := GenDisplaceFn(argsFloats[0], argsFloats[1], argsFloats[2])
	argsTime := text("Time")
	fmt.Printf("Displacement: %g\n",displacement(argsTime[0]))
}