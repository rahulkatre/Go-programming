package main

import "fmt"

var x int
var y string
var z bool
func main()  {
	fmt.Println(x,y,z)
	fmt.Printf("%T, %T, %T", x, y, z)
}
