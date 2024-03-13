// Program which allows the user to store a set of animals and request information.

package main

import (
	"fmt"
)

// declare a map which takes elements of type 'AnimalProps' (defined below); map is defined in main()
var animal_data = map[string]AnimalProps{}

// declare+define slices with the values
var animal_kinds = []string{"cow", "bird", "snake"}
var animal_actions = []string{"eat", "move", "speak"}

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// the Eat() method should print the animal’s food,
// the	Move() method should print the animal's locomotion, and
// the Speak() method should print the animal's spoken sound.
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Define three types Cow, Bird, and Snake.
// For each of these three types, define methods Eat(), Move(), and Speak() so that
// the types Cow, Bird, and Snake all satisfy the Animal interface.
type AnimalProps struct {
	name string
}
type Cow AnimalProps
type Bird AnimalProps
type Snake AnimalProps

// declare a slice for the animal records, init with size=0 and capacity=1
// data-type is (a pointer to) the interface(!), which is satisfied by the methods of the receiving actual data-types (kind of objects in Go)
var animal_records = make([]Animal, 0, 1)

// type Cow/Bird/Snake receive methods which satisfy the Animal interface
func (c Cow) Eat() {
	fmt.Println(c.name + " is a cow and eats grass.")
}
func (c Cow) Move() {
	fmt.Println(c.name + " is a cow and the locomotion method is walk.")
}
func (c Cow) Speak() {
	fmt.Println(c.name + " is a cow and the spoken sound is moo.")
}

func (b Bird) Eat() {
	fmt.Println(b.name + " is a bird eats worms.")
}
func (b Bird) Move() {
	fmt.Println(b.name + " is a bird and the locomotion method is fly.")
}
func (b Bird) Speak() {
	fmt.Println(b.name + " is a bird and the spoken sound is peep.")
}

func (s Snake) Eat() {
	fmt.Println(s.name + " is a snake eats mice.")
}
func (s Snake) Move() {
	fmt.Println(s.name + " is a snake and the locomotion method is slither.")
}
func (s Snake) Speak() {
	fmt.Println(s.name + " is a snake and the spoken sound is hsss.")
}

// check the existence of an item in a list
func itemInList(item string, list []string) bool {
	for index, _ := range list {
		if item == list[index] {
			return true
		}
	}
	return false
}

func getAnimalsName(a Animal) string {
	var name string
	name = ""
	// extract the name from the object
	if obj, ok := a.(Cow); ok {
		name = obj.name
	} else if obj, ok := a.(Bird); ok {
		name = obj.name
	} else {
		obj, _ := a.(Snake)
		name = obj.name
	}
	return name
}

func main() {
	// user input variables
	var input_command, input_name, input_argument string
	var input_count int
	var err error

	// hardcode animal information in a data structure
	// define the map
	//	animal_data[animal_kinds[0]] = AnimalProps{food: "grass", locomotion: "walk", noise: "moo"}
	//	animal_data[animal_kinds[1]] = AnimalProps{food: "worms", locomotion: "fly", noise: "peep"}
	//	animal_data[animal_kinds[2]] = AnimalProps{food: "mice", locomotion: "slither", noise: "hsss"}

	//
	// procedure
	//
	//	Each “newanimal” command must be a single line containing three strings.
	//	The first string is “newanimal”.
	//	The second string is an arbitrary string which will be the name of the new animal.
	//	The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
	//
	//	Each “query” command must be a single line containing 3 strings.
	//	The first string is “query”.
	//	The second string is the name of the animal.
	//	The third string is the name of the information requested about the animal, either “eat”, ‘move’, or “speak”.
	//	Your program should process each query command by printing out the requested data.

	fmt.Println("---------")
	fmt.Println("Create an animal record or request information about an existing one.")
	fmt.Println("\nCreate a record")
	fmt.Println("\tCommand: newanimal NAME KIND")
	fmt.Println("\tNAME => Give the animal a name.")
	fmt.Println("\tKIND => " + fmt.Sprint(animal_kinds))
	fmt.Println("\nQuery")
	fmt.Println("\tCommand: query NAME INFO")
	fmt.Println("\tNAME => The animal's name.")
	fmt.Println("\tINFO => " + fmt.Sprint(animal_actions) + "\n")
	fmt.Println("p \t list entries")
	fmt.Println("BLANK \t exit the program\n---------")

	for {
		// prompt user input
		fmt.Print("\nPlease create or request an animal record. \n>")
		input_count, err = fmt.Scanf("%s %s %s", &input_command, &input_name, &input_argument)

		// evaluate input
		if input_count == 0 {
			break
		}
		if input_count == 1 && input_command == "p" {
			var name string
			for _, item := range animal_records {
				name = getAnimalsName(item)
				fmt.Println(name)
			}
			if name == "" {
				fmt.Println("No entries yet.")
			}
			continue
		}
		if err != nil {
			fmt.Println("Please enter 3 values.")
			continue
		}
		if input_command != "newanimal" && input_command != "query" {
			fmt.Println("First value needs to be 'newanimal' or 'query'")
			continue
		}
		if input_command == "newanimal" && itemInList(input_argument, animal_kinds) == false {
			fmt.Println(input_argument + " is invalid. Choose from: " + fmt.Sprint(animal_kinds)) // print an array
			continue
		}
		if input_command == "query" && itemInList(input_argument, animal_actions) == false {
			fmt.Println(input_argument + " is invalid. Choose from: " + fmt.Sprint(animal_actions)) // print an array
			continue
		}

		// output

		if input_command == "newanimal" {
			// When the user creates an animal, create an object of the appropriate type.
			// Your program should process each newanimal command by creating the new animal and
			// printing “Created it!” on the screen.
			//			var a Animal
			switch input_argument {
			case animal_kinds[0]:
				var a Cow
				a.name = input_name
				animal_records = append(animal_records, a)
			case animal_kinds[1]:
				var a Bird
				a.name = input_name
				animal_records = append(animal_records, a)
			case animal_kinds[2]:
				var a Snake
				a.name = input_name
				animal_records = append(animal_records, a)
			}
			fmt.Println("Created it!")
		} else {
			//	Your program should call the appropriate method when the user issues a query command.
			var name string
			for _, item := range animal_records {
				// extract the name from the object
				name = getAnimalsName(item)
				if name == input_name {
					// calling the right function using the interface variable
					switch input_argument {
					case animal_actions[0]:
						item.Eat()
					case animal_actions[1]:
						item.Move()
					case animal_actions[2]:
						item.Speak()
					}
				}
			}
			if name == "" {
				fmt.Println("No entry for " + input_name + " yet.")
			}
		}
	}
}
