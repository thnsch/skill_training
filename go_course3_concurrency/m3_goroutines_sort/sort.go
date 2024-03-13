package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var input_integers_slice, a1, a2, a3, a4 []int

	// prompt and wait for user input
	fmt.Print("\n--- Sort integers ---\nPlease enter comma-separated integers, at least 4. E.g. 4,3,2,1 \n>")
	scanner.Scan()

	// split input into an int array using json unmarshal
	err := json.Unmarshal([]byte("["+scanner.Text()+"]"), &input_integers_slice)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Printf("%d", input_integers_slice[0])

	if len(input_integers_slice) < 4 {
		fmt.Printf("Expected at least 4 integer values, got %d.\n", len(input_integers_slice))
		os.Exit(1)
	}

	// default size of an array which is going to a gorouting
	sort_unit_len := len(input_integers_slice) / 4

	// create channels, call the goroutines to sort the sub-arrays
	c1 := make(chan []int)
	go sort_arrays(1, input_integers_slice[:sort_unit_len], c1)

	c2 := make(chan []int)
	go sort_arrays(2, input_integers_slice[sort_unit_len:sort_unit_len*2], c2)

	c3 := make(chan []int)
	go sort_arrays(3, input_integers_slice[sort_unit_len*2:sort_unit_len*3], c3)

	c4 := make(chan []int)
	go sort_arrays(4, input_integers_slice[sort_unit_len*3:], c4)

	// receive the sorted sub-arrays
	a1 = <-c1
	a2 = <-c2
	a3 = <-c3
	a4 = <-c4

	fmt.Printf("Goroutine 1 sent: %d\n", a1)
	fmt.Printf("Goroutine 2 sent: %d\n", a2)
	fmt.Printf("Goroutine 3 sent: %d\n", a3)
	fmt.Printf("Goroutine 4 sent: %d\n", a4)

	// merge the sub-arrays
	a1 = append(a1, a2...)
	a1 = append(a1, a3...)
	a1 = append(a1, a4...)

	// sort the whole array
	slices.Sort(a1)
	fmt.Printf("=> Sorted array: %d\n", a1)
}

func sort_arrays(caller_number int, values []int, c chan []int) {
	// print the argument containing the sub-array
	fmt.Printf("Goroutine %d will sort: %d\n", caller_number, values)
	slices.Sort(values)
	c <- values
}
