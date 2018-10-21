// Sample app to test crash a5 for xor-gate/goexif2.
package main

import (
	"fmt"
	"os"

	"github.com/xor-gate/goexif2/exif"
)

func main() {
	f, err := os.Open("crashers\\a59a2ad5701156b88c6a132e1340fe006f67280c")
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
