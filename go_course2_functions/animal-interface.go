/*
Test the program by running it and issuing three newanimal commands followed by three query commands. Each new animal should involve a different animal type (cow, bird, snake), each with a different name. Each query should involve a different animal and a different property of the animal (eat, move, speak). Give 2 points for each query for which the program provides the correct response.

Examine the code. If the code contains an interface type called Animal, which is a struct containing three fields, all of which are strings, then give another 2 points. If the program contains three types – Cow, Bird, and Snake – which all satisfy the Animal interface, give another 2 points.

Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, ">", to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a "newanimal" command or a "query" command.

Each "newanimal" command must be a single line containing three strings. The first string is "newanimal". The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either "cow", "bird", or "snake".  Your program should process each newanimal command by creating the new animal and printing "Created it!" on the screen.

Each "query" command must be a single line containing 3 strings. The first string is "query". The second string is the name of the animal. The third string is the name of the information requested about the animal, either "eat", "move", or "speak". Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface` should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

newanimal john cow
query john eat
query john2 move
"cow", "bird", or "snake"
"eat", "move", or "speak"
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const command_new = "newanimal"
const command_query = "query"
const info_eat = "eat"
const info_move = "move"
const info_speak = "speak"
const type_cow = "cow"
const type_bird = "bird"
const type_snake = "snake"

type Animal interface {
	Eat()
	Move()
	Speak()
}

// ========== Cow ==========
type Cow struct {
	Food       string
	Locomotion string
	Noise      string
}

func (cow *Cow) Eat() {
	fmt.Println(cow.Food)
}
func (cow *Cow) Move() {
	fmt.Println(cow.Locomotion)
}
func (cow *Cow) Speak() {
	fmt.Println(cow.Noise)
}

// ========== Bird ==========
type Bird struct {
	Food       string
	Locomotion string
	Noise      string
}

func (cow *Bird) Eat() {
	fmt.Println(cow.Food)
}
func (cow *Bird) Move() {
	fmt.Println(cow.Locomotion)
}
func (cow *Bird) Speak() {
	fmt.Println(cow.Noise)
}

// ========== Snake ==========
type Snake struct {
	Food       string
	Locomotion string
	Noise      string
}

func (cow *Snake) Eat() {
	fmt.Println(cow.Food)
}
func (cow *Snake) Move() {
	fmt.Println(cow.Locomotion)
}
func (cow *Snake) Speak() {
	fmt.Println(cow.Noise)
}

// ========== New Animal ==========
func createAnimal(name string, typeAnimal string, animalMap *map[string]Animal) {
	switch typeAnimal {
	case type_cow:
		cow := Cow{"grass", "walk", "moo"}
		(*animalMap)[name] = &cow
	case type_bird:
		bird := Bird{"worms", "fly", "peep"}
		(*animalMap)[name] = &bird
	case type_snake:
		snake := Snake{"mice", "slither", "hsss"}
		(*animalMap)[name] = &snake
	default:
		fmt.Println("Invalid animal type (cow, bird, snake)")
		return
	}
	fmt.Println("Created it!")
}

// ========== Query Animal ==========
func queryAnimal(name string, info string, animalMap *map[string]Animal) {
	if (*animalMap)[name] == nil {
		fmt.Println("Animal's name not found")
		return
	}

	switch info {
	case info_eat:
		(*animalMap)[name].Eat()
	case info_move:
		(*animalMap)[name].Move()
	case info_speak:
		(*animalMap)[name].Speak()
	default:
		fmt.Println("Invalid animal info (eat, move, speak)")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	animalMap := make(map[string]Animal)
	for {
		fmt.Print("> ")
		line, _, _ := reader.ReadLine()
		lineString := strings.TrimSpace(string(line))
		lineStringLower := strings.ToLower(lineString)
		word := strings.Fields(lineStringLower)

		if len(word) == 3 {
			switch word[0] {
			case command_new:
				createAnimal(word[1], word[2], &animalMap)
			case command_query:
				queryAnimal(word[1], word[2], &animalMap)
			default:
				fmt.Println("Invalid command (newanimal, query)")
			}

		} else if strings.ToLower(lineString) == "x" {
			break
		} else {
			fmt.Println("Invalid command")
		}
	}
}
