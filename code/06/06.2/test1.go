// Small program to test panic when calling Uint32(nil).
package main

import (
	"encoding/binary"
)

func main() {
	_ = binary.BigEndian.Uint32(nil)
	// _ = binary.BigEndian.Uint32([]byte(nil))

}
