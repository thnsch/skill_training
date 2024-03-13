package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create InputReader
	ir := bufio.NewReader(os.Stdin)

	// Ask for input
	fmt.Print("Please enter a name: ")
	name, _ := ir.ReadString('\n')
	name = strings.TrimSuffix(name, "\n") // remove trailing LF, which was expected by ReadString

	fmt.Print("Please enter an address: ")
	address, _ := ir.ReadString('\n')
	address = strings.TrimSuffix(address, "\n") // remove trailing LF, which was expected by ReadString

	// add to map; keys: name, address
	personMap := map[string]string{"name": name, "address": address}

	// use Marshal() to create JSON from the map
	personJson, _ := json.Marshal(personMap)

	// print JSON object
	fmt.Printf("JSON representation of map content: \n%s \n", personJson)
}
