// Sample app to test crash a5 for xor-gate/goexif2.
package main

import (
	"fmt"
	"os"

	"github.com/xor-gate/goexif2/exif"
)

func main() {
	f, err := os.Open("crashers\\49dfc363adbbe5aac9c2f8afbb0591c3ef1de2c3")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = exif.Decode(f)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("no err")
}
