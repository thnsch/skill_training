package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		nameLen = 20
	)

	// define name struct, two fields 'fname' 'lname'
	type Name struct {
		fname string
		lname string
	}

	nameSlice := make([]Name, 0)

	// ask for filename; content: "fname lname" per line
	stdinScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the name of the text file: ")
	stdinScanner.Scan()

	// get a file handle
	fileHandle, err := os.Open(stdinScanner.Text())
	if err != nil {
		fmt.Printf("Error opening file: %s \n%v \n", stdinScanner.Text(), err)
		os.Exit(1)
	}
	defer fileHandle.Close()

	// read each line, and store in name slice
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() { // read next line
		fullname := strings.Split(fileScanner.Text(), " ") // split the line into an array, separator is the space

		// store fullname in a name-struct
		// with max. length for each fname and lname
		var nameInstance Name
		if len(fullname[0]) > nameLen {
			nameInstance.fname = fullname[0][:nameLen]
		} else {
			nameInstance.fname = fullname[0]
		}
		if len(fullname[1]) > nameLen {
			nameInstance.lname = fullname[1][:nameLen]
		} else {
			nameInstance.lname = fullname[1]
		}

		// add to slice of type name
		nameSlice = append(nameSlice, nameInstance)
	}
	// check for file-scan errors
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file: %s \n%v \n", stdinScanner.Text(), err)
		os.Exit(1)
	}

	// iterate over slice, and print first and last names
	for i := range nameSlice {
		fmt.Printf("First name: %s \nLast name : %s \n\n", nameSlice[i].fname, nameSlice[i].lname)
	}
}

//bytesRead, _ := ioutil.ReadFile("something.txt")
//fileContent := string(bytesRead)
//lines := strings.Split(fileContent, "\n")
