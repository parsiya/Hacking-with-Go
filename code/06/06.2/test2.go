// Small program to investigate a panic in iprange for invalid masks.
package main

import "github.com/malfunkt/iprange"

func main() {
	_ = Fuzz([]byte("0.0.0.0/40"))
}

func Fuzz(data []byte) int {
	_, err := iprange.ParseList(string(data))
	if err != nil {
		return 0
	}
	return 1
}
