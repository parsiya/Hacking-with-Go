package main

import "fmt"

func main() {

	// Create a slice pointing to an int array
	s1 := make([]int, 5)

	fmt.Println(s1) // [0 0 0 0 0]

	for i := 0; i < len(s1); i++ {
		s1[i] = i
	}

	fmt.Println(s1) // [0 1 2 3 4]

	s2 := make([]int, 3)

	for i := 0; i < len(s2); i++ {
		s2[i] = i
	}

	fmt.Println(s2) // [0 1 2]

	s3 := append(s1, s2...)

	fmt.Println(s3) // [0 1 2 3 4 0 1 2]
}
