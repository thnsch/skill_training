package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// empty integer slice len 3
	sli := make([]int, 0, 3)
	myInput := ""
	myInt := 0

	for {
		// prompt for interger
		fmt.Print("Enter an integer value (or X to quit): ")
		fmt.Scan(&myInput)

		// end condition, capital X
		if myInput == "X" {
			break
		}

		// convert to integer
		myInt, _ = strconv.Atoi(myInput)

		// add to slice
		sli = append(sli, myInt)

		// sort ascending
		sort.Slice(sli, func(i, j int) bool {
			return sli[i] < sli[j]
		})

		// print the sorted slice
		fmt.Printf("Sorted slice: %d \n", sli)
	}

	fmt.Println("Goodbye!")
}
