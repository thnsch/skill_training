package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BubbleSort(toBeSorted []int) {

	// loop until the slice is sorted
	for !sort.SliceIsSorted(toBeSorted, func(i, j int) bool { return toBeSorted[i] < toBeSorted[j] }) {
		// loop over the slice
		for index, _ := range toBeSorted {
			// don't call swap for the last slice item, there is no i+1 left
			if index+1 == len(toBeSorted) {
				// fmt.Println(index, " , ", len(toBeSorted))
				break
			}
			swap(toBeSorted, index) // call swap for every index
		}
	}
}

func swap(toBeSorted []int, i int) {

	var tmp int

	if toBeSorted[i] > toBeSorted[i+1] {
		tmp = toBeSorted[i]             // mind the left (bigger) one
		toBeSorted[i] = toBeSorted[i+1] // pull the smaller one to the left
		toBeSorted[i+1] = tmp           // put the bigger one to the right
	}
}

func main() {

	const amountOfIntegers = 10

	var userInput []string
	intValues := make([]int, 0, amountOfIntegers) // declare an empty slice: 0 = size, amountOfIntegers = capacity
	stdinScanner := bufio.NewScanner(os.Stdin)
	//fmt.Println(len(intValues), cap(intValues))

	// wait for the correct amount of input values
	for {
		// ask for the integers
		fmt.Printf("Please enter %d blank-separated integers and hit enter, e.g. 1 2 3 4 ...: \n", amountOfIntegers)
		stdinScanner.Scan()

		userInput = strings.Fields(stdinScanner.Text()) // trims, and splits by removing all separating blanks, returns []string

		// evaluate amount of values
		if len(userInput) != amountOfIntegers {
			fmt.Printf("Found %d value(s), expected %d. \n\n", len(userInput), amountOfIntegers)
		} else {
			break
		}
	}

	// loop over userInput
	for _, value := range userInput {

		// convert item to integer
		myInt, err := strconv.Atoi(value)

		if err == nil {
			// append item to slice
			intValues = append(intValues, myInt)
		} else {
			fmt.Printf("Error %s \n", err)
			os.Exit(1)
		}
	}

	// sort the slice items
	BubbleSort(intValues)

	// print slice
	fmt.Printf("The sorted slice: %d\n", intValues)

	// print integers in one line, sorted asc

}
