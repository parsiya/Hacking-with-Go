// Get int and pointer size.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var p *int
	var p2 *float32

	fmt.Printf("Size of int      : %d\n", unsafe.Sizeof(i))
	fmt.Printf("Size of *int     : %d\n", unsafe.Sizeof(p))
	fmt.Printf("Size of *float32 : %d\n", unsafe.Sizeof(p2))
}
