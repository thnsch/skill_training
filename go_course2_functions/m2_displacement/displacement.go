package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a float64, v_0 float64, s_0 float64) func(float64) float64 {
	f_displace := func(t float64) float64 {
		displace := 0.5*a*math.Pow(t, 2) + v_0*t + s_0
		return displace
	}

	return f_displace
}

func main() {

	var err error
	var accelaration, initial_velocity, initial_displacement, time, result float64

	stdinScanner := bufio.NewScanner(os.Stdin)

	//
	// prompt for initial values
	//

	// prompt for acceleration, and evaluate input
	for {
		fmt.Println("Please enter accelaration: ")
		stdinScanner.Scan()
		// convert input-string to float64
		accelaration, err = strconv.ParseFloat(strings.TrimSpace(stdinScanner.Text()), 64)
		if err != nil {
			fmt.Println("\n" + err.Error()) // err.Error() function returns type string, needed to concat
		} else {
			break
		}
	}

	// prompt for initial_velocity, and evaluate input
	for {
		fmt.Println("Please enter initial_velocity: ")
		stdinScanner.Scan()
		// convert input-string to float64
		initial_velocity, err = strconv.ParseFloat(strings.TrimSpace(stdinScanner.Text()), 64)
		if err != nil {
			fmt.Println("\n" + err.Error())
		} else {
			break
		}
	}

	// prompt for initial_displacement, and evaluate input
	for {
		fmt.Println("Please enter initial_displacement: ")
		stdinScanner.Scan()
		// convert string to float64
		initial_displacement, err = strconv.ParseFloat(strings.TrimSpace(stdinScanner.Text()), 64)
		if err != nil {
			fmt.Println("\n" + err.Error())
		} else {
			break
		}
	}

	//fmt.Println("Value for acc is " + strconv.FormatFloat(accelaration, 'f', 2, 64))
	//fmt.Println("Value for ini_velo is " + strconv.FormatFloat(initial_velocity, 'f', 2, 64))
	//fmt.Println("Value for ini_displ is " + strconv.FormatFloat(initial_displacement, 'f', 2, 64))

	// generate the function, which is based on the entered initial values
	fmt.Println("\nGenerate the function ...")
	fn := GenDisplaceFn(accelaration, initial_velocity, initial_displacement)

	for {
		// prompt for the time
		fmt.Println("\nPlease enter a value for time (press x to quit) : ")
		stdinScanner.Scan()

		if stdinScanner.Text() == "x" {
			os.Exit(0)
		}

		// convert string to float
		time, err = strconv.ParseFloat(strings.TrimSpace(stdinScanner.Text()), 64)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// call the generated function
		result = fn(time)

		// convert float to string, and print result
		fmt.Println("--> Displacement after " + strconv.FormatFloat(time, 'f', 2, 64) + " seconds is " + strconv.FormatFloat(result, 'f', 2, 64))
	}
}
