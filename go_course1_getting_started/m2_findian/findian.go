
package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"unicode/utf8"
)

func main() {

	fmt.Println("Enter a text that contains the characters 'i', 'a' and 'n'.")
	ir := bufio.NewReader(os.Stdin)	// "InputReader" object, reads from cli including spaces
	s, _ := ir.ReadString('\n')
	s = strings.ToLower(s)

	if strings.Index(s, "i") == 0 &&	// starts with i
	   strings.ContainsRune(s, 'a') &&	// contains a
	   strings.LastIndex(s, "n") == utf8.RuneCountInString(s) - 2 {	// ends with n
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}