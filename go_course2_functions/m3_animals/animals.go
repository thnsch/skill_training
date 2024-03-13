// Program which allows the user to get information about a predefined set of animals.
// Three animals are predefined, cow, bird, and snake.
// Each animal can eat, move, and speak.
// The user can issue a request to find out one of three things about an animal.

package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println("Food: " + a.food)
}

func (a Animal) Move() {
	fmt.Println("Locomotion: " + a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("Noise: " + a.noise)
}

// check the existence of an item in a string
func itemInList(item string, list []string) bool {
	for index, _ := range list {
		if item == list[index] {
			return true
		}
	}
	return false
}

func main() {

	// user input variables
	var input_animal, input_info string
	var input_count int
	var err error

	// declare a map which takes elements of type 'Animal' (declared above)
	animals_data := map[string]Animal{}

	// define slices with the values
	animals := []string{"cow", "bird", "snake"}
	properties := []string{"eat", "move", "speak"}

	// hardcode animal information in a data structure
	// define the map
	animals_data[animals[0]] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals_data[animals[1]] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals_data[animals[2]] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	//
	// procedure
	//
	fmt.Println("---------")
	fmt.Println("Request information about animals. Choose from")
	fmt.Println("Animals:    " + fmt.Sprint(animals))
	fmt.Println("Properties: " + fmt.Sprint(properties) + "\n---------")
	fmt.Println("Empty input will exit the program.\n---------")

	for {
		// prompt user input
		fmt.Print("\nPlease enter the animal and property (e.g. 'cow eat'). \n>")
		input_count, err = fmt.Scanf("%s %s", &input_animal, &input_info)

		// evaluate input
		if input_count == 0 {
			break
		}
		if err != nil {
			fmt.Println("Please enter 2 values.")
			continue
		}
		if itemInList(input_animal, animals) == false {
			fmt.Println(input_animal + " is invalid. Choose from: " + fmt.Sprint(animals)) // print an array
			continue
		}
		if itemInList(input_info, properties) == false {
			fmt.Println(input_info + " is invalid. Choose from: " + fmt.Sprint(properties)) // print an array
			continue
		}

		// output
		fmt.Println("Animal: " + input_animal)

		switch input_info {
		case properties[0]: //eat
			animals_data[input_animal].Eat()
		case properties[1]: //move
			animals_data[input_animal].Move()
		case properties[2]: //speak
			animals_data[input_animal].Speak()
		default:
		}
	}
}
