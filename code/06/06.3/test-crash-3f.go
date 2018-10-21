// Sample app to test crash 3f for xor-gate/goexif2.
package main

import (
	"fmt"
	"os"

	"github.com/xor-gate/goexif2/exif"
)

func main() {
	f, err := os.Open("crashers\\3f5b7d448a0791f5739fa0a2371bb2492b64f835")
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
