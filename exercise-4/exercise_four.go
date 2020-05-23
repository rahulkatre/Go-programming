package main

import "fmt"

type IZ int
var x IZ
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x=42
	fmt.Println(x)
	//y = int(x)
	x = IZ(y)
	fmt.Println(x)
	fmt.Printf("%T\n", IZ(y))
}
