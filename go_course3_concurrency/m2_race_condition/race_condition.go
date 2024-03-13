package main

import (
	"fmt"
	"time"
)

func print_a() {
	fmt.Print(" a")
}

func print_b() {
	fmt.Print(" b")
}

func main() {
	rounds := 5

	for i := 1; i <= rounds; i++ {
		fmt.Printf("---- Run %d: ", i)
		time.Sleep(time.Second / 2)

		go print_a()
		go print_b()

		time.Sleep(time.Second / 2)
		fmt.Println()
	}

	time.Sleep(time.Second / 2)
}

// Explanation:
// 
// The functions print_a and print_b are called as goroutines (go ...).
// They are executed independently from each other.
// 
// With every run, it depends on which routine gets CPU time first.
// Therefore, the program will print "a b" or "b a".
// 
// It's not predictable which routine will be processed first.