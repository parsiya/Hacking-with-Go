package main

import "fmt"

// Tuple type
type Tuple struct {
	A, B int
}

// Should not change the value of the object as it works on a copy of it
func (x Tuple) ModifyTupleValue() {
	x.A = 2
	x.B = 2
}

// Should change the value of the object
func (x *Tuple) ModifyTuplePointer() {
	x.A = 3
	x.B = 3
}

type IntSlice []int

func (x IntSlice) PrintSlice() {
	fmt.Println(x)
}

// Modifies the IntSlice although it's by value
func (x IntSlice) ModifySliceValue() {
	x[0] = 1
}

// Modifies the IntSlice
func (x *IntSlice) ModifySlicePointer() {
	(*x)[0] = 2
}

func main() {

	tup := Tuple{1, 1}

	tup.ModifyTupleValue()
	fmt.Println(tup) // {1 1} - Does not change

	tup.ModifyTuplePointer()
	fmt.Println(tup) // {3 3} - Modified by pointer receiver

	var slice1 IntSlice = make([]int, 5)
	slice1.PrintSlice() // [0 0 0 0 0]

	slice1.ModifySliceValue()
	slice1.PrintSlice() // [1 0 0 0 0]

	slice1.ModifySlicePointer()
	slice1.PrintSlice() // [2 0 0 0 0]
}
