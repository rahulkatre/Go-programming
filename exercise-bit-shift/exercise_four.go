package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d, %b, %x\n", x, x, x)
	// Shifting x by 1 bit and assigning to variable y
	y := x << 1
	fmt.Printf("%d, %b, %x", y, y, y)
}

