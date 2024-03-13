package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func limite(buffer string) string {
	if len(buffer) > 20 {
		return string(buffer[0:20])
	}
	return buffer
}

type person struct {
	fname string
	lname string
}

var p = fmt.Println

func main() {

	//SE DECLARA EL SLICE DE PERSONAS DE TAMAÑO y CON O VALORES
	s := make([]person, 0, 0)

	//SE DECLARA VARIABLE ARCHIVO
	var archivo string
	p("Enter file name:")
	fmt.Scan(&archivo)

	readFile, err := os.Open(archivo)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range fileTextLines {
		splitNames := strings.Split(eachline, " ")
		nombre := splitNames[0]
		apellido := splitNames[1]
		p := person{fname: limite(nombre), lname: limite(apellido)}
		s = append(s, p)
	}
	for i := 0; i < len(s); i++ {
		p("fname: "+s[i].fname, " - lname: "+s[i].lname)
	}
}
