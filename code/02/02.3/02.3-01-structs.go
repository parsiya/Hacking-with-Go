package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

func main() {
	// Make an instance
	studentOne := Student{"Ender", "Wiggin"}

	// Now we can access fields
	fmt.Println(studentOne.FirstName)

	// We can just assign fields using names, anything not assigned will be
	// initialized with "zero" as we have seen before
	studentTwo := Student{FirstName: "Petra"}

	// We will print "{Petra }" notice the space after Petra which is supposed
	// to be the delimiter between the fields, LastName is nil because it is not
	// given a value
	fmt.Println(studentTwo)

	// Can also make a pointer to a struct
	p := &studentOne

	// Now instead of *p.LastName (doesn't work) we can just use p.LastName
	// fmt.Println((*p).LastName) will not work with error message: invalid indirect of p (type Student)
	fmt.Println(p.LastName)

	// Which is the same as
	fmt.Println(studentOne.LastName)

	// We can just create a pointer out of the blue
	p2 := &Student{"Hercule", "Poirot"}
	fmt.Println(p2)
}
