package main

import "fmt"

type intMap map[int]int

// Create a Stringer for this map type
func (i intMap) String() string {

	var s string
	s += fmt.Sprintf("Map type %T\n", i)
	s += fmt.Sprintf("Length: %d\n", len(i))

	// Iterate through all key/value pairs
	for k, v := range i {
		s += fmt.Sprintf("[%v] = %v\n", k, v)
	}
	return s
}

func main() {

	// Create a map
	map1 := make(intMap)

	// Add key/value pairs
	map1[0] = 10
	map1[5] = 20

	// Print map - Stringer will be called
	fmt.Println(map1)
	// Map type main.intMap
	// Length: 2
	// [0] = 10
	// [5] = 20

	// Delete a key/value pair
	delete(map1, 0)

	fmt.Println(map1)
	// Map type main.intMap
	// Length: 1
	// [5] = 20

	// Create a map on the spot using members
	map2 := map[string]string{"key1": "value1", "key2": "value2"}

	fmt.Println(map2)
	// map[key1:value1 key2:value2]
}
