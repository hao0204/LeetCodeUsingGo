package main

import (
	"fmt"
	"strings"
	// "encoding/hex"
)
func main() {
	data := []byte{49, 46, 48}
	str := string(data[:])
	// a := new string("1.0")
	fmt.Println(strings.EqualFold(str, "1.0"))
}