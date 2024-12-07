package main

import (
	"fmt"
)

func main() {
	b := 1 << 30
	fmt.Printf("%d bytes\n", b)
	fmt.Printf("%b\n", b)
}
