package main

import "fmt"

func main() {

	// Create an array of strings with 3 members
	characters := [3]string{"Ender", "Petra", "Mazer"}

	// Last index is exclusive
	// allMembers []string := characters[0:3]
	var allMembers []string = characters[0:3]
	fmt.Println("All members", allMembers)

	var lastTwo []string = characters[1:3]
	fmt.Println("Last two members", lastTwo)

	// Replace Mazer with Bean
	fmt.Println("Replacing Mazer with Bean")
	allMembers[2] = "Bean"

	fmt.Println("All members after Bean swap", characters)

	fmt.Println("Last two members after Bean swap", lastTwo)
}
