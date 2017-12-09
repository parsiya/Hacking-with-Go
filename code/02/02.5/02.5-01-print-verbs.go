package main

import "fmt"

type myType struct {
	field1 int
	field2 string
	field3 float64
}

func main() {

	// int
	fmt.Println("int:")
	int1 := 123
	fmt.Printf("%v\n", int1)     // 123
	fmt.Printf("%d\n", int1)     // 123
	fmt.Printf("|%6d|\n", int1)  // |   123|
	fmt.Printf("|%-6d|\n", int1) // |123   |
	fmt.Printf("%T\n", int1)     // int
	fmt.Printf("%x\n", int1)     // 7b
	fmt.Printf("%b\n", int1)     // 1111011
	fmt.Printf("%e\n", int1)     // %!e(int=123)
	fmt.Printf("%c\n", int1)     // { - 0x7B = 123
	fmt.Println()

	// float
	fmt.Println("float:")
	float1 := 1234.56
	fmt.Printf("%f\n", float1)       // 1234.560000
	fmt.Printf("|%3.2f|\n", float1)  // |1234.56|
	fmt.Printf("|%-3.2f|\n", float1) // |1234.56|
	fmt.Printf("%e\n", float1)       // 1.234560e+03
	fmt.Printf("%E\n", float1)       // 1.234560E+03
	fmt.Println()

	// string
	fmt.Println("string:")
	string1 := "Petra"
	fmt.Printf("%s\n", string1)      // Petra
	fmt.Printf("|%10s|\n", string1)  // |     Petra|
	fmt.Printf("|%-10s|\n", string1) // |Petra     |
	fmt.Printf("%T\n", string1)      // string
	fmt.Println()

	// boolean
	fmt.Println("boolean:")
	boolean1 := true
	fmt.Printf("%t\n", boolean1) // true
	fmt.Printf("%T\n", boolean1) // bool
	fmt.Println()

	// struct type
	fmt.Println("struct:")
	struct1 := myType{10, "Ender", -10.2}
	fmt.Printf("%v\n", struct1)  // {10 Ender -10.2}
	fmt.Printf("%+v\n", struct1) // {field1:10 field2:Ender field3:-10.2}
	fmt.Printf("%#v\n", struct1) // main.myType{field1:10, field2:"Ender", field3:-10.2}
	fmt.Printf("%T\n", struct1)  // main.myType
}
